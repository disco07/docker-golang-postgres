package main

import (
	"net/http"
	"time"
)

type post struct {
	ID          int
	Title       string
	Slug        string
	Summary     string
	Content     string
	PublishedAt time.Time
}

func getPosts(w http.ResponseWriter, r *http.ResponseWriter) {

}

func main() {
	http.HandleFunc("/posts", getPosts)
	http.ListenAndServe(":8000", nil)
}
