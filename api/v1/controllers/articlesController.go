package controllers

import (
	"build-sections/api/v1/config/db"
	"build-sections/api/v1/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type articlesController struct{}

var Article articlesController

func (s articlesController) ListAll(rw http.ResponseWriter, req *http.Request) {

	articles := []models.Article{}
	db.DB.Find(&articles)
	respondJSON(rw, http.StatusOK, articles)
}

func (s articlesController) Create(rw http.ResponseWriter, req *http.Request) {
	article := models.Article{}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&article); err != nil {
		respondError(rw, http.StatusBadRequest, err.Error())
		return
	}
	defer req.Body.Close()

	if err := db.DB.Save(&article).Error; err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(rw, http.StatusCreated, article)
}

func (s articlesController) Show(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])
	article := getArticleOr404(id, rw, req)
	if article == nil {
		return
	}
	respondJSON(rw, http.StatusOK, article)
}

func getArticleOr404(id int, rw http.ResponseWriter, req *http.Request) *models.Article {
	article := models.Article{}
	if err := db.DB.First(&article, id).Error; err != nil {
		respondError(rw, http.StatusNotFound, err.Error())
		return nil
	}
	return &article
}
