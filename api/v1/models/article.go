package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string `valid:"alphanum,required"`
	Author  string `valid:"alphanum,required"`
	Content string `valid:"alphanum,required"`
}

type Response struct {
	Status  int
	Message string
}
