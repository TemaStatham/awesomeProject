package app

import (
	"awesomeProject/config"
	"awesomeProject/internal/repository/postgres"
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
		Host:     cfg.DBConfig.Host,
		Port:     cfg.DBConfig.Port,
		Username: cfg.DBConfig.Username,
		Password: cfg.DBConfig.Password,
		DBName:   cfg.DBConfig.DBName,
		SSLMode:  cfg.DBConfig.SSLMode,
	})
	if err != nil {
		panic(err)
	}
}
