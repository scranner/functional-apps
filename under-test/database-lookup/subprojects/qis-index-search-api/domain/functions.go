package domain

import (
	"encoding/json"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
)

type Result struct {
	Url string `url:"json"`
	Title string `title:"json"`
}

type SearchResult struct {
	Results []Result `results:"json"`
}

var (
	queensIndex = redisearch.NewClient(os.Getenv("QUEENS_INDEX_URL") + ":6379", "SearchIndex")
	schema = redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("url", redisearch.TextFieldOptions{ Weight:   40 })).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{ Weight:   40 })).
		AddField(redisearch.NewTextFieldOptions("keywords", redisearch.TextFieldOptions{ Weight:   20 })).
		AddField(redisearch.NewTextFieldOptions("h1", redisearch.TextFieldOptions{ Weight:   20 })).
		AddField(redisearch.NewTextFieldOptions("h2", redisearch.TextFieldOptions{ Weight:   10 })).
		AddField(redisearch.NewTextFieldOptions("h3", redisearch.TextFieldOptions{ Weight:   5 })).
		AddField(redisearch.NewTextFieldOptions("h4", redisearch.TextFieldOptions{ Weight:   5 })).
		AddField(redisearch.NewTextFieldOptions("h5", redisearch.TextFieldOptions{ Weight:   5 })).
		AddField(redisearch.NewTextFieldOptions("h6", redisearch.TextFieldOptions{ Weight:   5 }))
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
	searchTerm :=  r.URL.Query().Get("searchTerm")

	if searchTerm == "" {
		http.Error(w, "Invalid Parameters", 400)
		return
	}

	docs, total, err := queensIndex.Search(redisearch.NewQuery(searchTerm).
		Limit(0, 2).
		SetReturnFields("title"))

	if err != nil {
		http.Error(w, "Lookup Error", 500)
		return
	}

	resultsToReturn := SearchResult {
		make([]Result, total),
	}

	for i, doc := range docs {
		resultsToReturn.Results[i] = Result{
			Url:   doc.Id,
			Title: fmt.Sprintf("%v", doc.Properties["title"]),
		}
	}

	textToWrite, err1 := json.Marshal(resultsToReturn)

	if err1 != nil {
		_, _ = fmt.Fprint(w,  err1.Error())
		return
	}

	_, _ = w.Write(textToWrite)
}
