package main

import (
	"net/http"
	"regexp"
	"time"
)

var reUUID *regexp.Regexp

func init() {
	// 910bf50a-31fd-4820-a4e9-f4115aad529a
	reUUID = regexp.MustCompile(`(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})`)
}

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
	// /job/910bf50a-31fd-4820-a4e9-f4115aad529a
	match := reUUID.FindStringSubmatch(r.URL.Path)
	if len(match) != 2 {
		http.Redirect(w, r, "/jobs?error=not-found", http.StatusSeeOther)
		return
	}

	result := db.First(&j, "uuid = ?", match[1])

	if result.Error != nil {
		http.Redirect(w, r, "/jobs?error=not-found", http.StatusSeeOther)
		return
	}

	action := r.FormValue("action")

	switch action {
	case "del":
		db.Delete(&j)
		http.Redirect(w, r, "/jobs", http.StatusSeeOther)
		return
	case "do":
		j.Due = time.Now()
		db.Save(&j)
	}

	tpls.ExecuteTemplate(w, "job.gohtml", j)
}
