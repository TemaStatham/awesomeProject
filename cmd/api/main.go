package main

import (
	"awesomeProject/internal/app"
)

func main() {
	a := app.NewApp()
	a.Run()
}

// go run ./cmd/api/main.go --config="./config/config.yaml"
//docker-compose up -d --build
