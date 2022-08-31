package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type config struct {
	port int
	dsn  string
}

type post struct {
	ID          int
	Title       string
	Slug        string
	Summary     string
	Content     string
	PublishedAt time.Time
}

func getPosts(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	var cfg = config{
		port: 8000,
		dsn:  "postgres://blog:blog@localhost/blog?sslmode=disable",
	}

	db, err := open(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/posts", getPosts)
	http.ListenAndServe(fmt.Sprintf(":%v", cfg.port), nil)
}

func open(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
