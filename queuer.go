package main

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

type queuer struct {
	db      *gorm.DB
	running bool
	options queueOptions
}

func newQueue(db *gorm.DB, opt queueOptions) *queuer {
	if db == nil {
		log.Fatalln("DB must not be nil!")
	}
	return &queuer{
		db:      db,
		running: false,
		options: opt,
	}
}

func (q *queuer) digestForeverEvery(d time.Duration) {
	if q.running {
		log.Println("Queue already running")
		return
	}
	q.running = true

	// this will block the goroutine
	for _ = range time.Tick(d) {
		q.tick()
	}
}

func (q *queuer) tick() {
	log.Println("Ticking!")
	chn := make(chan job)
	mu := sync.Mutex{}
	i := 0
	for _, j := range q.getJobs() {
		log.Println("NAME:", j.Name)
		i++
		go q.do(j, chn)
	}

	if i == 0 {
		return
	}

	for c := range chn {
		mu.Lock()
		i--
		db.Save(c)
		if i <= 0 {
			log.Println("Closing chan...")
			close(chn)
		}
		mu.Unlock()
	}
}

func (q *queuer) do(j job, chn chan job) {
	log.Printf("Executing job '%s' (%s)\n", j.Name, j.UUID)
	res, err := http.Post(
		j.Endpoint,
		j.ContentType,
		strings.NewReader(j.Payload),
	)

	if res != nil {
		log.Printf(
			"Job sent to the endpoint (%s), received the HTTP Status Code: %d (this job requires: %d)\n",
			j.Endpoint,
			res.StatusCode,
			j.HTTPOkStatus,
		)
	}

	if err != nil || res == nil || res.StatusCode != j.HTTPOkStatus {
		j.Retries++
		log.Printf("Job errored. Times errored %d/%d\n",
			j.Retries, q.options.maxRetries)
		if j.Retries >= q.options.maxRetries {
			j.Status = statusFailed
		} else {
			j.Status = statusRequeued
		}
	} else {
		log.Printf("Job done after %d retry/ies\n", j.Retries)
		j.Status = statusDone
	}

	// send back the modified job
	chn <- j
}

func (q queuer) getJobs() []job {
	var xj []job
	db.Order(
		"due ASC",
	).Limit(
		q.options.maxConcurrent,
	).Find(
		&xj, "due <= ? AND status IN (?) AND retries < ?",
		time.Now(),
		[]string{statusRequeued, statusQueued},
		q.options.maxRetries,
	)
	log.Printf("Found %d jobs!\n", len(xj))
	for _, j := range xj {
		log.Println(j.Name)
	}
	return xj
}
