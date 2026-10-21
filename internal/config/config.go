package config

import (
	"errors"
	"log"
	"net/url"
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
		u = strings.TrimSpace(u)
		err := validateURL(u)
		if err != nil {
			log.Fatalf("invalid target %q: %v", u, err)
		}
		targets = append(targets, u)
	}

	return Config{
		Port:          port,
		CheckInterval: d,
		Targets:       targets,
	}
}

// validateURL проверяет, что строка — корректный HTTP(S)-адрес.
func validateURL(raw string) error {
	u, err := url.Parse(raw) // функция из пакета net/url
	if err != nil {
		return err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("invalid Scheme")
	}
	if u.Host == "" {
		return errors.New("invalid Host")
	}
	return nil
}
