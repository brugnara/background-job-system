package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
		//
		log.Println("Something called me @", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
