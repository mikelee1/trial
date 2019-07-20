package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/op/go-logging"
)

func main() {
	var log = logging.MustGetLogger("prometheus")
	// expose prometheus metrics接口
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8085", nil))
}