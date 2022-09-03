package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFindAllPost(t *testing.T) {
	app := newApp()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	posts, err := app.findAllPost(ctx)

	if err != nil {
		t.Error("Get All Posts failed.")
	}

	if len(posts) == 0 {
		t.Error("Posts did not return any values.")
	}

}

func TestGetPosts(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req := httptest.NewRequest("GET", "/posts", nil)

	app := newApp()
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getPosts)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var posts []post
	err := json.NewDecoder(rr.Body).Decode(&posts)
	if err != nil {
		t.Error(err.Error())
		t.Error("Error retreiving list of posts.")
	}

	if len(posts) == 0 {
		t.Error("Error retreiving list of posts.")
	}
}

func newApp() apps {
	cfg := config{
		port: 8000,
		dsn:  "postgres://app:app@localhost/app?sslmode=disable",
	}

	db, _ := open(cfg)
	return apps{
		DB: db,
	}
}
