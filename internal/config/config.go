package config

import "os"

// Config — все настройки приложения в одном месте.
type Config struct {
	// Port — адрес, который слушает HTTP-сервер (например, ":8080")
	Port string
}

// New читает окружение и возвращает Config с разумными дефолтами.
func New() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	return Config{
		Port: port,
	}
}
