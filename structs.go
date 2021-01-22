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
	Name         string
	Retries      int
	Payload      string
	Due          time.Time
	Endpoint     string
	ContentType  string
	HTTPOkStatus int
	Done         bool
}
