package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

// Config - структура конфига
type Config struct {
	Env          string         `yaml:"env" env-default:"local"`
	StoragePaths string         `yaml:"storage_paths" env-required:"true"`
	DBConfig     PostgresConfig `yaml:"postgres" env-required:"true"`
}

type PostgresConfig struct {
	Port     string `yaml:"port" env-default:"5436"`
	Host     string `yaml:"host" env-default:"localhost"`
	Username string `yaml:"username" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-default:"5436"`
	Password string `yaml:"password" env-required:"true"`
}

// MustLoad получает структуру конфига
func MustLoad() *Config {
	path := fetchConfigFlags()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}

	return &cfg
}

// fetchConfigFlags получает путь до конфига либо из флага командной строки либо через переменную окружения
func fetchConfigFlags() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
