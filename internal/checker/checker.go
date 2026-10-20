package checker

import (
	"net/http"
	"time"
	// 1. импорты: net/http, time, ваши config (если понадобится) и metrics
)

type Checker struct {
	// client — ОДИН на все проверки. Новый клиент на каждый запрос —
	// антипаттерн: он не переиспользует TCP-соединения и жрёт ресурсы.
	client *http.Client
	urls   []string
	ticker *time.Ticker
}

func New(urls []string, interval time.Duration) *Checker {
	// 2. вернуть &Checker{
	//     client: &http.Client{Timeout: 10 * time.Second},
	//     urls:   urls,
	//     ticker: time.NewTicker(interval),
	// }
	return &Checker{}
}

// Run — блокирующий бесконечный цикл. Его вызовут в отдельной горутине.
func (c *Checker) Run() {
	// 3. for range c.ticker.C {  <- тикер сам «щёлкает» каждые interval
	//        for _, u := range c.urls {
	//            c.check(u)
	//        }
	//    }
}

func (c *Checker) check(rawURL string) {
	// 4. start := time.Now()

	// 5. resp, err := c.client.Get(rawURL)
	//    duration := time.Since(start) — засекаем ВСЕГДА, до ветвления

	// 6. metrics.CheckDuration.WithLabelValues(rawURL).Observe(duration.Seconds())
	//    (этоHistogram: у него метод Observe, а не Set/Inc)

	// 7. если err != nil — соединение не состоялось:
	//        ChecksTotal.WithLabelValues(rawURL, "failure").Inc()
	//        LastStatusCode.WithLabelValues(rawURL).Set(0)
	//        return
	//    иначе:
	//        defer resp.Body.Close()  <- сразу после проверки err! Иначе соединения утекут.
	//        Подумайте: код 200 — success, а 404 и 500? Запрос-то Состоялся.
	//        Успех = 2xx–3xx. Определите outcome по resp.StatusCode
	//        ChecksTotal.WithLabelValues(rawURL, outcome).Inc()
	//        LastStatusCode.WithLabelValues(rawURL).Set(float64(resp.StatusCode))
}
