package main

import (
	"fmt"
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
		fmt.Fprintln(w, "Hello world!")
		fmt.Fprintf(w, "Host: %s\n", host)
		fmt.Fprintf(w, "User-Agent: %s\n", r.UserAgent())
		fmt.Fprintf(w, "Source Host: %s\n", r.Host)
		fmt.Fprintf(w, "Remote Address: %s\n", r.RemoteAddr)
		fmt.Fprintf(w, "X-Forward-Address: %s\n", r.Header.Get("X-Forwarded-For"))
		fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
