package app

import (
	"awesomeProject/config"
	clickhousedb "awesomeProject/pkg/database/clickhouse"
	"awesomeProject/pkg/database/postgres"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	cfg := config.MustLoad()

	// инит кликхаус

	_, err := postgres.NewPostgresDB(postgres.Config{
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

	_, err = clickhousedb.NewClickhouseDB(clickhousedb.Config{
		Host:     "",
		Port:     "",
		Database: "",
		Username: "",
		Password: "",
		Client:   "",
		Version:  "",
	})
	if err != nil {
		panic(err)
	}
}
