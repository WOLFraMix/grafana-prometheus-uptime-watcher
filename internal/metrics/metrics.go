package metrics

import "github.com/prometheus/client_golang/prometheus"

// ChecksTotal — сколько проверок сделано, с лейблами url и outcome.
// outcome будет принимать значения "success" | "failure" —
// так одна метрика покрывает оба случая, фильтрация в Grafana по лейблу.
var ChecksTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "watcher_checks_total",
		Help: "Total number of checks performed.",
	},
	[]string{"url", "outcome"},
)

// LastStatusCode — код последнего ответа, с лейблом url.
var LastStatusCode = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "watcher_last_status_code",
		Help: "HTTP status code of the last check.",
	},
	[]string{"url"},
)

// CheckDuration — время выполнения проверки.
var CheckDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "watcher_check_duration_seconds",
		Help:    "Duration of a single check.",
		Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2, 5},
	},
	[]string{"url"},
)

// Init регистрирует все метрики в дефолтном реестре.
// Без регистрации promhttp их не отдаст!
func Init() {
	prometheus.MustRegister(ChecksTotal)
	prometheus.MustRegister(LastStatusCode)
	prometheus.MustRegister(CheckDuration)
}
