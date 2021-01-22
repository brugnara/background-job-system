package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

// index handles HTTP requests
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		log.Println("Handling post..")
		indexAction(w, r)
		return
	}
	//
	w.Header().Set("Content-Type", "text/html")
	err := tpls.ExecuteTemplate(w, "index.gohtml", struct {
		Error string
		Done  string
	}{
		Error: r.FormValue("error"),
		Done:  r.FormValue("done"),
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func indexAction(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	ep := r.FormValue("ep")
	ct := r.FormValue("ct")
	pl := r.FormValue("pl")
	// 2021-01-22T12:41
	d := r.FormValue("due") + ":00Z"
	due, err := time.Parse(time.RFC3339, d)
	if err != nil {
		log.Println("Invalid date:", d)
		http.Redirect(w, r, "/?error=invalid-date", http.StatusSeeOther)
		return
	}
	cd, err := strconv.Atoi(r.FormValue("cd"))
	if err != nil {
		cd = 200
	}
	j := job{
		UUID:         fmt.Sprintf("%s", uuid.NewV4()),
		Name:         name,
		Endpoint:     ep,
		ContentType:  ct,
		Payload:      pl,
		Due:          due,
		Status:       statusQueued,
		HTTPOkStatus: cd,
	}
	log.Printf("Creating job: %s (%s)\n", name, j.UUID)
	db.Create(&j)
	http.Redirect(w, r, fmt.Sprintf("/?done=%s", j.UUID), http.StatusSeeOther)
}
