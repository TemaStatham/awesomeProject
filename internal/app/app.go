package app

import (
	"awesomeProject/config"
	"awesomeProject/internal/handler"
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

	_, err = clickhousedb.NewClickhouseDB(&clickhousedb.Config{
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

	_, err = redisdb.NewRedisDB(&redisdb.Config{
		Addr:     cfg.RConfig.Addr,
		Password: cfg.RConfig.Password,
		DB:       cfg.RConfig.DB,
	})
	if err != nil {
		panic(err)
	}

	hand := handler.NewHandler(l)

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
		l.Info("error occured on db connection close: %s", err.Error())
		return
	}

	return
}
