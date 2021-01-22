package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var q *queuer
var db *gorm.DB

const (
	dbDrv = "sqlite3"
	dbURL = "./database.db"
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connected")
	// preparing db
	db.AutoMigrate(&job{})
	log.Println("DB ready")
	//
	tmpJob := job{
		Name:     "tmp job",
		Payload:  "pippo",
		Endpoint: "http://localhost:8080/pippo",
		Due:      time.Now().Add(-5 * time.Minute),
	}
	db.Create(&tmpJob)
	db.Create(&job{
		Name: "ciccio",
		Due:  time.Now().Add(+5 * time.Minute),
	})
	//
	q = newQueue(db, queueOptions{
		maxRetries:    5,
		maxConcurrent: 10,
	})
	//
	// cleanup stuff
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		db.Delete(&tmpJob, 1)
		log.Println("Exiting now...")
		os.Exit(0)
	}()
}

func main() {
	go serveHTTP(":8081")
	q.digestForeverEvery(5 * time.Second)
}
