package config

import (
	"database/sql"
	"log"
	"os"

	"build-sections/api/v1/config/db"
)

func SetupEnv(goEnv string) (env string) {

	var err error

	switch goEnv {
	case "development":
		os.Setenv("development", "development")
		log.Println(os.Getenv("development"))
		db.DBCon, err = sql.Open("postgres", "user=postgres password=password dbname=build_sections_development sslmode=disable")
		if err != nil {
			panic(err)
		}
	}
	return os.Getenv("development")
}
