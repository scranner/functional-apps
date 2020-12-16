package main

import (
	"github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.hal.davecutting.uk/jmccartney13/qis-ad-submission-api/domain"
	"log"
	"net/http"
)

func main() {

	domain.Init()

	router := chi.NewRouter()

	prometheusMiddleware := chiprometheus.NewMiddleware("ad-api")
	router.Use(prometheusMiddleware)
	domain.InstallCors(router)

	router.Handle("/metrics", promhttp.Handler())
	router.Post("/", domain.RootHandler)
	router.Get("/live", domain.StatusHandler)
	router.Get("/ready", domain.StatusHandler)

	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}
