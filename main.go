package main

import (
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

func init() {

}

func main() {
	http.HandleFunc("/posts", getPosts)
	http.ListenAndServe(":8000", nil)
}
