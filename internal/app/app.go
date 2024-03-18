package app

import (
	"awesomeProject/config"
	clickhousedb "awesomeProject/pkg/database/clickhouse"
	postgresdb "awesomeProject/pkg/database/postgres"
	redisdb "awesomeProject/pkg/database/redis"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	cfg := config.MustLoad()

	_, err := postgresdb.NewPostgresDB(&postgresdb.Config{
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
		Host:     cfg.ChConfig.Host,
		Port:     cfg.ChConfig.Port,
		Database: cfg.ChConfig.Database,
		Username: cfg.ChConfig.Username,
		Password: cfg.ChConfig.Password,
		Client:   cfg.ChConfig.Client,
		Version:  cfg.ChConfig.Version,
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
}
