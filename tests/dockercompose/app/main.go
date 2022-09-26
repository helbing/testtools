package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil { // nolint:gosec
		log.Fatal(err)
	}
}
