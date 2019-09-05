package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/tapas-sharma/prometheusMetrics/restServer/dummy"
	_ "github.com/tapas-sharma/prometheusMetrics/restServer/dummy/prometheus"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://0.0.0.0:8080"
)

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func main() {
	fmt.Println("Starting prometheus metric rest service")
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
		service  = flag.String("service", "prometheus", "Pass the service to start")
	)

	logger := log.NewLogfmtLogger(os.Stderr)

	httpLogger := log.With(logger, "component", "http")
	svc, err := dummy.Get(*service)
	if err != nil {
		panic(err)
	}
	mux := mux.NewRouter().StrictSlash(false)
	dummy.MakeHandler(svc, httpLogger, mux)
	for k, v := range svc.GetCustomRoutes() {
		mux.Handle(k, v)
	}
	//http.Handle("/", mux)

	errs := make(chan error, 2)
	//start the http server
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening", "service", *service)
		errs <- http.ListenAndServe(*httpAddr, mux)
	}()
	//start the signal handle to handle ctrl + c gracefully
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
