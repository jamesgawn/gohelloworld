package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	host, err := os.Hostname()

	if err != nil {
		log.Panicf("An error has occured: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hello world!")
		log.Printf("Host: %s", host)
		log.Printf("User-Agent: %s", r.UserAgent())
		log.Printf("Source Host: %s", r.Host)
		log.Printf("Remote Address: %s", r.RemoteAddr)
		log.Printf("X-Forward-Address: %s", r.Header.Get("X-Forwarded-For"))
		log.Printf("Path: %s", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
