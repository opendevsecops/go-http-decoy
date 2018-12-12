package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func basicAuthHandler(w http.ResponseWriter, r *http.Request) {
}

func digestAuthHandler(w http.ResponseWriter, r *http.Request) {
}

func ntlmAuthHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	host := flag.String("host", "0.0.0.0", "Host address to listen to")
	port := flag.Int("port", 8080, "Port number to listen to")
	basicAuth := flag.String("basic-auth", "", "Force basic authentication on path")
	digestAuth := flag.String("digets-auth", "", "Force digest authentication on path")
	ntlmAuth := flag.String("ntlm-auth", "", "Force ntlm authentication on path")

	flag.Parse()

	if *basicAuth != "" {
		http.HandleFunc(*basicAuth, basicAuthHandler)
	}

	if *digestAuth != "" {
		http.HandleFunc(*digestAuth, digestAuthHandler)
	}

	if *ntlmAuth != "" {
		http.HandleFunc(*ntlmAuth, ntlmAuthHandler)
	}

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to listen for http connections on address %s:%d", *host, *port))

		return
	}
}
