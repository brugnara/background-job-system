package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpls *template.Template

func init() {
	// tpls stuff
	tpls = template.Must(template.ParseGlob("./tpls/*.gohtml"))
	// handlers
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := tpls.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func serveHTTP(s string) {
	http.ListenAndServe(s, nil)
}
