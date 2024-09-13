package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "handler"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCount.With(prometheus.Labels{"method": r.Method, "handler": "/"})
		w.Write([]byte("Hello, Prometheus!"))
	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
