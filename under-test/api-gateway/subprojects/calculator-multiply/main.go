package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")
	result, err := Multiply(x, y)

	log.Println(err)

	if err == nil {
		_, _ = fmt.Fprintf(w, "{\"result\":\"" + fmt.Sprintf("%.0f", result) + "\"}")
	} else {
		http.Error(w, "Invalid Parameter", 400)
	}

	log.Println("GET /")

}

func live(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "OK")
}

func ready(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "OK")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", root)
	router.HandleFunc("/live", live)
	router.HandleFunc("/ready", ready)
	log.Println("Server now available on :80")
	log.Fatal(http.ListenAndServe(":80", router))
}