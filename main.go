package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// var limiter *rate.Limiter

func main() {
	corsObj := handlers.AllowedOrigins([]string{"*"})
	// limiter = rate.NewLimiter(config.RateLimit, config.RateLimitBurst)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/init", ksInit)
	router.HandleFunc("/show", ksShow)
	router.HandleFunc("/generate", ksGenerate)
	router.HandleFunc("/depget", index)
	router.HandleFunc("/env", index)
	log.Printf("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ksInit(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "stubs/init.json")
}

func ksShow(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "stubs/show.json")
}
func ksGenerate(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "stubs/generate.json")
}
