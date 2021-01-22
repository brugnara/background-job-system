package main

import (
	"net/http"
)

func jobsIndex(w http.ResponseWriter, r *http.Request) {
	var xj []job
	db.Order(
		"due ASC",
	).Offset(0).Limit(20).Find(&xj)
	//
	tpls.ExecuteTemplate(w, "jobs.gohtml", xj)
}

func jobIndex(w http.ResponseWriter, r *http.Request) {
	var j job
	tpls.ExecuteTemplate(w, "job.gohtml", j)
}
