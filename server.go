package main

import (
	"net/http"
)

type formInput struct {
	Type        string
	Name        string
	Label       string
	Placeholder string
	Value       string
}

func init() {
	// handlers
	http.HandleFunc("/", index)
	http.HandleFunc("/jobs", jobsIndex)
	http.HandleFunc("/job/", jobIndex)
	http.HandleFunc("/css/bulma.min.css", bulmaCSS)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func serveHTTP(s string) {
	http.ListenAndServe(s, nil)
}
