package app

import (
	"awesomeProject/config"
	"awesomeProject/internal/handler"
	postgresrepos "awesomeProject/internal/repository/postgres"
	service "awesomeProject/internal/service"
	clickhousedb "awesomeProject/pkg/database/clickhouse"
	postgresdb "awesomeProject/pkg/database/postgres"
	redisdb "awesomeProject/pkg/database/redis"
	"awesomeProject/pkg/logger"
	"awesomeProject/pkg/server"
	"context"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	cfg := config.MustLoad()

	l := logger.SetupLogger(cfg.Env)

	pg, err := postgresdb.NewPostgresDB(&postgresdb.Config{
		Host:     cfg.PgConfig.Host,
		Port:     cfg.PgConfig.Port,
		Username: cfg.PgConfig.Username,
		Password: cfg.PgConfig.Password,
		DBName:   cfg.PgConfig.DBName,
		SSLMode:  cfg.PgConfig.SSLMode,
	})
	if err != nil {
		panic(err)
	}

	ch, err := clickhousedb.NewClickhouseDB(&clickhousedb.Config{
		HttpPort:   cfg.ChConfig.HttpPort,
		NativePort: cfg.ChConfig.NativePort,
		Addr:       cfg.ChConfig.Addr,
		Database:   cfg.ChConfig.Database,
		Username:   cfg.ChConfig.Username,
		Password:   cfg.ChConfig.Password,
	})
	if err != nil {
		panic(err)
	}

	r, err := redisdb.NewRedisDB(&redisdb.Config{
		Addr:     cfg.RConfig.Addr,
		Password: cfg.RConfig.Password,
		DB:       cfg.RConfig.DB,
	})
	if err != nil {
		panic(err)
	}

	repos := postgresrepos.NewPostgres(pg, l)
	serv := service.NewService(l, repos, repos, repos, repos)
	hand := handler.NewHandler(l, serv, serv, serv, serv)

	srv := new(server.Server)
	go func() {
		if err := srv.Run("8000", hand.Init()); err != nil {
			l.Error("error occured while running http server: %s", err.Error())
		}
	}()

	l.Info("Application started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	l.Info("Application Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		l.Info("error occured on server shutting down: %s", err.Error())
		return
	}

	if err := pg.Close(); err != nil {
		l.Info("error occured on pg db connection close: %s", err.Error())
		return
	}

	if err := r.Close(); err != nil {
		l.Info("error occured on redis db connection close: %s", err.Error())
		return
	}

	if err := ch.Close(); err != nil {
		l.Info("error occured on clickhouse db connection close: %s", err.Error())
		return
	}

	return
}
