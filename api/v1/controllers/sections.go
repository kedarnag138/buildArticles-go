package controllers

import (
	"net/http"

	_ "github.com/lib/pq"
)

type sections struct{}

var Section sections

func (s sections) Create(rw http.ResponseWriter, req *http.Request) {

}
