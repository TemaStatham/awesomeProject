package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisDB(cfg *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(ping)

	return client, nil
}
