package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/satori/go.uuid"
)

type formInput struct {
	Type        string
	Name        string
	Label       string
	Placeholder string
	Value       string
}

var tpls *template.Template

func init() {
	// tpls stuff
	tpls = template.Must(
		template.New(
			"").Funcs(
			template.FuncMap{
				"randID":   randID,
				"toStruct": toStruct,
			},
		).ParseGlob("./tpls/*.gohtml"))
	// handlers
	http.HandleFunc("/", index)
	http.HandleFunc("/jobs", jobsIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

// randID generates an unique id for html tags
func randID() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}

// toStruct converts a `..|..|..|..` string, into a formInput{}
func toStruct(s string) formInput {
	tmp := strings.Split(s, "|")
	if len(tmp) != 5 {
		return formInput{}
	}
	return formInput{
		tmp[0],
		tmp[1],
		tmp[2],
		tmp[3],
		tmp[4],
	}
}

func serveHTTP(s string) {
	http.ListenAndServe(s, nil)
}
