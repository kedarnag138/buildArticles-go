package controllers

import (
	"build-sections/api/v1/config/db"
	"build-sections/api/v1/controllers"
	"build-sections/api/v1/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
)

type articlesController struct{}

var Article articlesController

func TestArticleCreate(t *testing.T) {
	var err error
	db.DB, err = gorm.Open("postgres", "user=postgres password=password dbname=build_articles_test sslmode=disable")
	db.DB.AutoMigrate(&models.Article{})
	if err != nil {
		panic(err)
	}
	inputJson := strings.NewReader(`{"title": "Sample", "author": "", "content": "sdfdggfdgdfg"}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/v1/sections", inputJson)

	controllers.Article.Create(w, r)
}

func TestListAllArticles(t *testing.T) {
	var err error
	db.DB, err = gorm.Open("postgres", "user=postgres password=password dbname=build_articles_test sslmode=disable")
	db.DB.AutoMigrate(&models.Article{})
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/articles", nil)

	controllers.Article.ListAll(w, r)
}

func TestGetArticleById(t *testing.T) {
	var err error
	db.DB, err = gorm.Open("postgres", "user=postgres password=password dbname=build_articles_test sslmode=disable")
	db.DB.AutoMigrate(&models.Article{})
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/articles/1", nil)

	controllers.Article.Show(w, r)
}
