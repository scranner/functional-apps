package main

import (
	"github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.hal.davecutting.uk/jmccartney13/qis-crawler/domain"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	domain.ConfigureLogging()

	domain.Init()

	router := chi.NewRouter()

	prometheusMiddleware := chiprometheus.NewMiddleware("crawler")
	router.Use(prometheusMiddleware)

	router.Handle("/metrics", promhttp.Handler())
	router.Get("/live", domain.StatusHandler)
	router.Get("/ready", domain.StatusHandler)

	go func() {
		if err := http.ListenAndServe(":80", router); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		loopTimer, err := strconv.Atoi(os.Getenv("CRAWL_DELAY")); if err != nil {
			loopTimer = 60
		}

		for {
			url, err := domain.PopQueue(); if err != nil || !domain.IsUrl(url) {
				log.Print("Failed to pop url from queue")
				time.Sleep(time.Second*10)
				continue
			}

			log.Printf("Scraping %s", url)

			result, err1 := domain.Scrape(url)

			if err1 != nil {
				log.Printf("Error scraping: %s\n%s", url, err.Error())
				time.Sleep(time.Second*20)
				continue
			}

			processedScrape := domain.ProcessScrape(result)
			err2 := domain.IndexAndStore(processedScrape)

			if err2 != nil {
				log.Printf("Error storing scrape: %s\n%s", url, err.Error())
				time.Sleep(time.Second*20)
				continue
			}

			log.Printf("Scrape complete")
			time.Sleep(time.Second * time.Duration(loopTimer))
		}
	}()

	select {}
}
