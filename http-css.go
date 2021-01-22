package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func bulmaCSS(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./public/bulma.min.css")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "not found")
		return
	}
	defer f.Close()
	w.Header().Set("content-type", "text/css")
	io.Copy(w, f)
}
