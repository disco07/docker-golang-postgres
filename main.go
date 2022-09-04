package main

import (
	"context"
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
	w.WriteHeader(status)

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Write(js)

	return nil
}

func (a *apps) findAllPost(ctx context.Context) ([]*post, error) {
	query := fmt.Sprintf(`SELECT * FROM post`)
	rows, err := a.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*post

	for rows.Next() {
		var post post
		err := rows.Scan(&post.ID, &post.Title, &post.Slug, &post.Summary, &post.Content, &post.PublishedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (a *apps) getPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := a.findAllPost(r.Context())
	if err != nil {
		JSON(w, http.StatusBadRequest, err)
	}

	JSON(w, http.StatusOK, posts)
}

func main() {
	var cfg = config{
		port: 8000,
		dsn:  "postgres://app:app@postgres-container/app?sslmode=disable",
	}

	db, err := open(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := apps{db}

	http.HandleFunc("/posts", app.getPosts)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.port), nil))
}

func open(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
