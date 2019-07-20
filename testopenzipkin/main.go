package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	tracer2 "myproj/try/testopenzipkin/tracer"
	"github.com/openzipkin/zipkin-go"
	"context"
)

var	tracer *zipkin.Tracer
var err error
func main() {
	tracer, err = tracer2.NewTracer()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/foo", FooHandler)
	r.Use(zipkinhttp.NewServerMiddleware(
		tracer,
		zipkinhttp.SpanName("request")), // name for request span
	)
	log.Fatal(http.ListenAndServe(":1080", r))
}

func CheckDB(ctx context.Context)  {
	span,ctx := tracer.StartSpanFromContext(ctx,"CheckDB",[]zipkin.SpanOption{}...)
	defer span.Finish()
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	CheckDB(ctx)

	span,ctx := tracer.StartSpanFromContext(ctx,"test",[]zipkin.SpanOption{}...)
	defer span.Finish()



	span,ctx = tracer.StartSpanFromContext(ctx,"test1",[]zipkin.SpanOption{}...)
	defer span.Finish()
	w.WriteHeader(http.StatusOK)
}