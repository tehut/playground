package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

func main() {
	limiter = rate.NewLimiter(config.RateLimit, config.RateLimitBurst)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/init", index)
	router.HandleFunc("/show", index)
	router.HandleFunc("/generate", index)
	router.HandleFunc("/depget", index)
	router.HandleFunc("/env", index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
