package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string
	Author  string
	Content string
}

type Response struct {
	Status  int
	Message string
	Data    uint
}
