package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "wip")
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "wip")
	})

	http.ListenAndServe(":8080", nil)
}
