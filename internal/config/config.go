package config

import (
	"log"
	"os"
	"strings"
	"time"
)

// Config — все настройки приложения в одном месте.
type Config struct {
	Port string
	// CheckInterval — как часто пинговать; уже разобранный time.Duration
	CheckInterval time.Duration
	// Targets — список URL из переменной TARGETS, разделённые запятой
	Targets []string
}

// New читает окружение и возвращает Config с разумными дефолтами.
func New() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	raw := os.Getenv("CHECK_INTERVAL")
	if raw == "" {
		raw = "15s"
	}
	d, err := time.ParseDuration(raw)
	if err != nil {
		log.Println("invalid CHECK_INTERVAL, using default 15s:", err)
		d = 15 * time.Second
	}

	raw = os.Getenv("TARGETS")
	if raw == "" {
		log.Fatal("empty urls")
	}
	urls := strings.Split(raw, ",")
	targets := []string{}
	for _, u := range urls {
		targets = append(targets, strings.TrimSpace(u))
	}

	return Config{
		Port:          port,
		CheckInterval: d,
		Targets:       targets,
	}
}
