package controllers

import (
	"build-sections/api/v1/config/db"
	"build-sections/api/v1/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

type articlesController struct{}

var Article articlesController

func (s articlesController) Create(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	var params models.Article

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &params)
	if err != nil {
		panic(err)
	}

	article := models.Article{Title: params.Title, Author: params.Author, Content: params.Content}

	response := db.DB.Create(&article).Scan(&article)

	if response != nil {
		b, err := json.Marshal(models.Response{
			Status:  201,
			Message: "success",
			Data:    article.ID,
		})
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(b)
	}
}
