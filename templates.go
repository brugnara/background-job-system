package main

import (
	"fmt"
	"html/template"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpls *template.Template

func init() {
	// tpls stuff
	tpls = template.Must(
		template.New(
			"").Funcs(
			template.FuncMap{
				"randID":   randID,
				"toStruct": toStruct,
				"toHero":   toHero,
			},
		).ParseGlob("./tpls/*.gohtml"))
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

func toHero(s string) hero {
	ret := hero{}
	tmp := strings.Split(s, "|")
	if len(tmp) > 0 {
		ret.Title = tmp[0]
	}
	if len(tmp) > 1 {
		ret.Subtitle = tmp[1]
	}
	return ret
}
