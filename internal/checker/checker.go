package checker

import (
	"log"
	"net/http"
	"time"

	"github.com/WOLFraMix/grafana-prometheus-uptime-watcher/internal/metrics"
)

type Checker struct {
	// client — ОДИН на все проверки.
	// Новый клиент на каждый запрос —
	// антипаттерн: он не переиспользует TCP-соединения и жрёт ресурсы.
	client *http.Client
	urls   []string
	ticker *time.Ticker
}

func New(urls []string, interval time.Duration) *Checker {
	return &Checker{
		client: &http.Client{Timeout: 10 * time.Second},
		urls:   urls,
		ticker: time.NewTicker(interval),
	}
}

// Run — блокирующий бесконечный цикл. Его вызовут в отдельной горутине.
func (c *Checker) Run() {
	for range c.ticker.C {
		for _, u := range c.urls {
			c.check(u)
		}
	}
}

func (c *Checker) check(rawURL string) {
	start := time.Now()

	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		// если запрос не смог родиться
		// — считаем провалом и выходим
		log.Println("bad request:", err)
		metrics.ChecksTotal.WithLabelValues(rawURL, "failure").Inc()
		metrics.LastStatusCode.WithLabelValues(rawURL).Set(0)
		return
	}
	req.Header.Set("User-Agent", "uptime-watcher/1.0")

	resp, err := c.client.Do(req)
	if err != nil {
		// соединение не состоялось: таймаут, DNS, отказ
		metrics.ChecksTotal.WithLabelValues(rawURL, "failure").Inc()
		metrics.LastStatusCode.WithLabelValues(rawURL).Set(0)
		return
	}
	defer resp.Body.Close()

	metrics.CheckDuration.WithLabelValues(rawURL).Observe(time.Since(start).Seconds())

	outcome := "failure"
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		outcome = "success"
	}
	metrics.ChecksTotal.WithLabelValues(rawURL, outcome).Inc()
	metrics.LastStatusCode.WithLabelValues(rawURL).Set(float64(resp.StatusCode))
}
