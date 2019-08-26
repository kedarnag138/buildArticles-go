package config

import (
	"build-sections/api/v1/config/db"
	"build-sections/api/v1/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func SetupEnv(goEnv string) (env string) {

	var err error

	switch goEnv {
	case "development":
		os.Setenv("development", "development")
		log.Println(os.Getenv("development"))
		db.DB, err = gorm.Open("postgres", "user=postgres password=password dbname=build_articles_development sslmode=disable")
		db.DB.AutoMigrate(&models.Article{})
		if err != nil {
			panic(err)
		}
	case "test":
		os.Setenv("test", "test")
		log.Println(os.Getenv("test"))
		db.DB, err = gorm.Open("postgres", "user=postgres password=password dbname=build_articles_test sslmode=disable")
		db.DB.AutoMigrate(&models.Article{})
		if err != nil {
			panic(err)
		}
	}
	return os.Getenv("development")
}
