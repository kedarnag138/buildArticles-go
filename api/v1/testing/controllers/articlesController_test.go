package controllers

import (
	"build-sections/api/v1/config"
	"build-sections/api/v1/controllers"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type articlesController struct{}

var Article articlesController

func TestCreateArticle(t *testing.T) {
	env := os.Getenv("env")
	config.SetupEnv(env)
	inputJSON := strings.NewReader(`{"title": "New Title", "author": "Constatine", "content": "Once upon a time in a far far away galaxy!"}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/v1/articles", inputJSON)
	r.Header.Set("Content-Type", "application/json")
	m := mux.NewRouter()
	m.HandleFunc("/api/v1/articles", controllers.Article.Create).Methods("POST")
	m.ServeHTTP(w, r)
}
