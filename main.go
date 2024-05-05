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

<<<<<<< HEAD
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "wip")
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "wip")
	})

=======
>>>>>>> 28b2bf3 (code)
	http.ListenAndServe(":8080", nil)
}
