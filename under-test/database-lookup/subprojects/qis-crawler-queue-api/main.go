package main

import (
	"github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.hal.davecutting.uk/jmccartney13/qis-crawler-queue-api/domain"
	"log"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	prometheusMiddleware := chiprometheus.NewMiddleware("queue-api")
	router.Use(prometheusMiddleware)

	router.Handle("/metrics", promhttp.Handler())
	router.Get("/live", domain.StatusHandler)
	router.Get("/ready", domain.StatusHandler)
	router.Post("/", domain.ProcessUrls)

	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}
