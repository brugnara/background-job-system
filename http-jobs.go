package main

import (
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var reUUID *regexp.Regexp

const pageSize = 20

func init() {
	// 910bf50a-31fd-4820-a4e9-f4115aad529a
	reUUID = regexp.MustCompile(`(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})`)
}

func jobsIndex(w http.ResponseWriter, r *http.Request) {
	var xj []job
	var count int64

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}

	db.Table("jobs").Count(&count)

	db.Order(
		"due ASC",
	).Offset((page - 1) * pageSize).Limit(pageSize).Find(&xj)

	pages := int(count / pageSize)
	if count%pageSize > 0 {
		pages++
	}

	// TODO: return a stuct with offset, limit and item count for
	// pagination!
	tpls.ExecuteTemplate(w, "jobs.gohtml", struct {
		Jobs  []job
		Page  int
		Pages int
		Count int64
	}{
		xj,
		page,
		pages,
		count,
	})
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
