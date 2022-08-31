package main

import (
	"database/sql"
	"encoding/json"
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

type apps struct {
	DB *sql.DB
}

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Write(js)

	return nil
}

func (a apps) getPosts(w http.ResponseWriter, r *http.Request) {

	query := fmt.Sprintf(`SELECT * FROM post`)
	rows, err := a.DB.QueryContext(r.Context(), query)
	if err != nil {
		JSON(w, http.StatusBadRequest, err)
	}
	defer rows.Close()

	var posts []*post

	for rows.Next() {
		var post post
		err := rows.Scan(&post.ID, &post.Title, &post.Slug, &post.Summary, &post.Content, &post.PublishedAt)
		if err != nil {
			JSON(w, http.StatusBadRequest, err)
		}
		posts = append(posts, &post)
	}

	JSON(w, http.StatusOK, posts)
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

	app := apps{db}

	http.HandleFunc("/posts", app.getPosts)
	http.ListenAndServe(fmt.Sprintf(":%v", cfg.port), nil)
}

func open(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
