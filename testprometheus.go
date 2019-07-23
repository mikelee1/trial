package main

import (
	"github.com/op/go-logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	var log = logging.MustGetLogger("prometheus")
	// expose prometheus metrics接口
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8085", nil))
}
