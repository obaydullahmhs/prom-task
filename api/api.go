package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func StartServer(port int) {
	prometheus.MustRegister(pingCounter)

	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
