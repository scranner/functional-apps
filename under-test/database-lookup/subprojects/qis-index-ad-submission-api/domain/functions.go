package domain

import (
	"encoding/json"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/rs/xid"
	"log"
	"net/http"
	"os"
	"time"
)

type advert struct {
	Title string `json:"title"`
	Advert string `json:"advert"`
	Keywords []string `json:"keywords"`
}

type adverts struct {
	Adverts []advert `json:"adverts"`
}

var (
	queensIndex = redisearch.NewClient(os.Getenv("QUEENS_INDEX_URL") + ":6379", "AdsIndex")
	schema = redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{ Weight:   100 })).
		AddField(redisearch.NewTextFieldOptions("keywords", redisearch.TextFieldOptions{ Weight:   50 })).
		AddField(redisearch.NewTextFieldOptions("advert", redisearch.TextFieldOptions{ Weight:   10 }))
)

func Init() {
	ConfigureLogging()
	_, err := queensIndex.Info()
	if err != nil {
		indexerr := queensIndex.CreateIndex(schema)
		if indexerr != nil {
			log.Fatal(indexerr.Error())
		}
	}
}

func InstallCors(router *chi.Mux) {
	cors := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
		AllowedMethods:         []string{"GET", "OPTIONS"},
		AllowedHeaders:         []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:         []string{"Link"},
		MaxAge:                 3599,
	})
	router.Use(cors.Handler)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w,  "{\"status\": \"OK\" } ")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var adverts adverts

	err := json.NewDecoder(r.Body).Decode(&adverts)

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	errors := make([]string, len(adverts.Adverts))

	for _, advert := range adverts.Adverts {
		doc := redisearch.NewDocument(xid.New().String(), 1.0).
			Set("title", advert.Title).
			Set("keywords", advert.Keywords).
			Set("advert", advert.Advert).
			Set("last_accessed", time.Now().Unix())

		if err := queensIndex.Index(doc); err != nil {
			errors = append(errors, err.Error())
		}

	}

	log.Printf("INFO: processed %d adverts. %d errors", len(adverts.Adverts), len(errors))
	textToWrite, err1 := json.Marshal(errors)

	if err1 != nil {
		_, _ = fmt.Fprint(w,  err1.Error())
		return
	}

	_, _ = w.Write(textToWrite)

}
