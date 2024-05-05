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

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "todo")
	})

	http.ListenAndServe(":8080", nil)
}
