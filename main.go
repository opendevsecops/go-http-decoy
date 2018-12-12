package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func LogHeader(header string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)

			if r.Header.Get(header) != "" {
				log.Println(r.Header.Get(header))
			}
		})
	}
}

func SetStatusText(code int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, http.StatusText(code), code)

			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	host := flag.String("host", "0.0.0.0", "Host address to listen to")
	port := flag.Int("port", 8080, "Port number to listen to")
	serverHeader := flag.String("server-header", "", "Set server header")
	basicAuth := flag.String("basic-auth", "", "Set basic auth realm")
	logCredentials := flag.Bool("log-credentials", false, "Log credentials")

	flag.Parse()

	r := chi.NewRouter()

	if *logCredentials {
		r.Use(LogHeader("Authorization"))
	}

	r.Use(middleware.Logger)

	if *serverHeader != "" {
		r.Use(middleware.SetHeader("Server", *serverHeader))
	}

	if *basicAuth != "" {
		r.Use(middleware.SetHeader("WWW-Authenticate", fmt.Sprintf("basic realm=%s", *basicAuth)))
		r.Use(SetStatusText(401))
	}

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), r)

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to listen for http connections on address %s:%d", *host, *port))

		return
	}
}
