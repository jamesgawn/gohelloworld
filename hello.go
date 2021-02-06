package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	host, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
		fmt.Fprintln(w, "Host: ", host)
		fmt.Fprintln(w, "User-Agent: ", r.UserAgent())
		fmt.Fprintln(w, "Source Host: ", r.Host)
		fmt.Fprintln(w, "Remote Address: ", r.RemoteAddr)
		fmt.Fprintln(w, "Path: ", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
