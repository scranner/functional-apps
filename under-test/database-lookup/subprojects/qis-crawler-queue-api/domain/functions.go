package domain

import (
	"encoding/json"
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	client, err = redis.Dial("tcp", os.Getenv("REDIS_URL") + ":6379")
	queue = redisqueue.New("worker-queue", client)
)

func Init() {
	if err != nil { panic("Failed to connect to redis") }
}

type urlList struct {
	Urls []string `json:"urls"`
}

type result struct {
	Results []string `json:"results"`
}

func (r result) Marshel() []byte {
	res, _ := json.Marshal(r)
	return res
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w,  "{\"status\": \"OK\" } ")
}

func ProcessUrls(w http.ResponseWriter, r *http.Request) {
	var urls urlList

	err := json.NewDecoder(r.Body).Decode(&urls)

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	results := make([]string, len(urls.Urls))

	for i, url := range urls.Urls {
		added, err := queue.Push(url)

		if err != nil {
			results[i] = err.Error()
			continue
		}

		results[i] = strconv.FormatBool(added)

	}

	log.Printf("INFO: processed %d urls", len(results))

	_, _ = w.Write(result { results }.Marshel())

}