package config

import (
	"os"
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
	return Config{
		Port: port,
	}

	// 2. CheckInterval:
	//    raw := os.Getenv("CHECK_INTERVAL"); если пусто — дефолт "15s"
	//    d, err := time.ParseDuration(raw)
	//    если err (кто-то написал "abc") — вернуть дефолт 15 секунд:
	//    падать из-за кривой опции не стоит, просто логируем и используем дефолт

	// 3. Targets:
	//    raw := os.Getenv("TARGETS"); если пусто — пустой срез []string{}
	//    strings.Split(raw, ",") — но каждый элемент подрезать strings.TrimSpace:
	//    "https://a.com, https://b.com" с пробелом после запятой не должно ломать парсинг
}
