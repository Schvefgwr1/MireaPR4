package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
	} `yaml:"database"`
	App struct {
		Port int `yaml:"port"`
	} `yaml:"app"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func LoadConfig(filename string) (*Config, error) {
	// Чтение данных из YAML-файла
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	// Переопределение значений из переменных окружения
	overrideFromEnv(&config)

	return &config, nil
}

func overrideFromEnv(config *Config) {
	// Функция-утилита для замены значений, если переменные окружения определены
	if env := os.Getenv("MARKET_DB_HOST"); env != "" {
		config.Database.Host = env
	}
	if env := os.Getenv("MARKET_DB_USER"); env != "" {
		config.Database.User = env
	}
	if env := os.Getenv("MARKET_DB_PASSWORD"); env != "" {
		config.Database.Password = env
	}
	if env := os.Getenv("MARKET_DB_NAME"); env != "" {
		config.Database.Name = env
	}
	if env := os.Getenv("MARKET_DB_PORT"); env != "" {
		if port, err := strconv.Atoi(env); err == nil {
			config.Database.Port = port
		}
	}
	if env := os.Getenv("MARKET_SERVER_PORT"); env != "" {
		if port, err := strconv.Atoi(env); err == nil {
			config.App.Port = port
		}
	}
}
