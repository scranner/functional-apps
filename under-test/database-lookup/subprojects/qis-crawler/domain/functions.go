package domain

import (
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	"github.com/gelembjuk/keyphrases"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
	"github.com/RediSearch/redisearch-go/redisearch"
)

var (
	client, redisConnectionError = redis.Dial("tcp", os.Getenv("REDIS_URL") + ":6379")
	queue = redisqueue.New("worker-queue", client)
	analyser = keyphrases.TextPhrases{ Language: "english" }
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

	// mention sitemap
)

type ScrapeResult struct {
	Url string
	Title string
	P []string
	H1 []string
	H2 []string
	H3 []string
	H4 []string
	H5 []string
	H6 []string
}

type ProcessedScrapeResult struct {
	Url string
	Title string
	Keywords []string
	H1 []string
	H2 []string
	H3 []string
	H4 []string
	H5 []string
	H6 []string
}

func Init() {
	if redisConnectionError != nil {
		log.Fatal("Failed to connect to redis")
	}
	if err := analyser.Init(); err != nil {
		log.Fatal(err.Error())
	}
	_, err := queensIndex.Info()

	if err != nil {
		indexerr := queensIndex.CreateIndex(schema)
		if indexerr != nil {
			log.Fatal(indexerr.Error())
		}
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w,  "{\"status\": \"OK\" } ")
}

func PopQueue() (string, error) {

	result, err := queue.Pop()

	if err != nil {
		return "", err
	}

	return result, nil
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func Scrape(url string) (ScrapeResult, error) {

	scraper := colly.NewCollector()

	var title string
	var pTag []string
	var h1Tag []string
	var h2Tag []string
	var h3Tag []string
	var h4Tag []string
	var h5Tag []string
	var h6Tag []string

	scraper.OnHTML("p", func(e *colly.HTMLElement) {
		pTag = append(pTag, processText(e.Text))
	})

	scraper.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	scraper.OnHTML("h1", func(e *colly.HTMLElement) {
		h1Tag = append(h1Tag, processText(e.Text))
	})

	scraper.OnHTML("h2", func(e *colly.HTMLElement) {
		h2Tag = append(h2Tag, processText(e.Text))
	})

	scraper.OnHTML("h3", func(e *colly.HTMLElement) {
		h3Tag = append(h3Tag, processText(e.Text))
	})

	scraper.OnHTML("h4", func(e *colly.HTMLElement) {
		h4Tag = append(h4Tag, processText(e.Text))
	})

	scraper.OnHTML("h5", func(e *colly.HTMLElement) {
		h5Tag = append(h5Tag, processText(e.Text))
	})

	scraper.OnHTML("h6", func(e *colly.HTMLElement) {
		h6Tag = append(h6Tag, processText(e.Text))
	})

	scraper.OnHTML("span.title", func(e *colly.HTMLElement) {
		h2Tag = append(h2Tag, processText(e.Text))
	})

	if err := scraper.Visit(url); err != nil {
		return ScrapeResult {}, err
	}

	return ScrapeResult{
		Url:   url,
		Title: title,
		P:     pTag,
		H1:    h1Tag,
		H2:    h2Tag,
		H3:    h3Tag,
		H4:    h4Tag,
		H5:    h5Tag,
		H6:    h6Tag,
	}, nil

}

func processText(text string) string {
	preWhitespace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	insideWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := preWhitespace.ReplaceAllString(text, "")
	return insideWhitespace.ReplaceAllString(final, " ")
}

func ProcessScrape(result ScrapeResult) ProcessedScrapeResult {

	keywords := analyser.GetKeyWords(strings.Join(result.P, " "))

	return ProcessedScrapeResult{
		Url:      result.Url,
		Title:    result.Title,
		Keywords: keywords,
		H1:		  result.H1,
		H2:		  result.H2,
		H3:		  result.H3,
		H4:		  result.H4,
		H5:		  result.H5,
		H6:		  result.H6,
	}
}

func IndexAndStore(result ProcessedScrapeResult) error {

	doc := redisearch.NewDocument(result.Url, 1.0).
		Set("title", result.Title).
		Set("url", result.Url).
		Set("keywords", result.Keywords).
		Set("last_accessed", time.Now().Unix()).
		Set("h1", result.H1).
		Set("h2", result.H2).
		Set("h3", result.H3).
		Set("h4", result.H4).
		Set("h5", result.H5).
		Set("h6", result.H6)

	if err := queensIndex.IndexOptions(redisearch.IndexingOptions{ Replace:  true }, doc); err != nil {
		return err
	}

	return nil
}