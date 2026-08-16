package main

import (
	"errors"
	"fmt"
	projetokorp "http-server-projeto-korp/internal/projeto_korp"
	"http-server-projeto-korp/internal/util"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	trace := util.CreateErrorContext("main")
	if err := webServer(); err != nil {
		panic(fmt.Sprintf("\n%s\n", trace.Apply(err)))
	}
}

func webServer() error {
	trace := util.CreateErrorContext("webServer")

	port, _port := util.EnvAsResult("HTTP_PORT")
	timeoutTime, _timeoutTime := util.EnvAsIntegerResult("TIMEOUT_TIME")
	projectName, _projectName := util.EnvAsResult("PROJECT_NAME")

	if err := errors.Join(_port, _timeoutTime, _projectName); err != nil {
		return trace.Apply(err)
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewBuildInfoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	router := mux.NewRouter().UseEncodedPath()

	router.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{})).Methods("GET")
	router.HandleFunc("/projeto-korp", projetokorp.GetProjetoKorp(projectName)).Methods("GET")

	fmt.Println("Starting HTTP SERVER PROJETO KORP at port ", port)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: time.Duration(timeoutTime) * time.Second,
		Handler:           router,
	}

	return trace.Apply(server.ListenAndServe())
}
