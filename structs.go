package main

import (
	"time"

	"gorm.io/gorm"
)

type queueOptions struct {
	maxRetries    int
	maxConcurrent int
}

type job struct {
	gorm.Model
	UUID         string
	Name         string
	Retries      int
	Payload      string
	Due          time.Time
	Endpoint     string
	ContentType  string
	HTTPOkStatus int
	Status       string
}

type hero struct {
	Title    string
	Subtitle string
}

const (
	statusDone     = "done"
	statusRequeued = "requeued"
	statusQueued   = "queued"
	statusFailed   = "failed"
)
