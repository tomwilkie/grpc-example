package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/felixge/httpsnoop"
	"github.com/prometheus/client_golang/prometheus"
)

var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "request_duration_seconds",
	Help:    "Time (in seconds) spent serving HTTP requests.",
	Buckets: prometheus.DefBuckets,
}, []string{"method", "route", "status_code"})

func init() {
	prometheus.MustRegister(requestDuration)
}

func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(h, w, r)
		requestDuration.WithLabelValues(r.Method, r.URL.Path,
			strconv.Itoa(m.Code)).Observe(m.Duration.Seconds())
	})
}

func server(addr string) {
	http.Handle("/metrics", prometheus.Handler())

	http.Handle("/greeter", wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := response{
			Message: "Hello " + req.Name,
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})))
	log.Fatalf("listen err: %v", http.ListenAndServe(addr, nil))
}
