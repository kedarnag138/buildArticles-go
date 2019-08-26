package main

import (
	"build-sections/api/v1/config"
	"build-sections/api/v1/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	getEnv := config.SetupEnv(os.Args[1])

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/sections", controllers.Section.Create).Methods("POST")

	// http.Handle("/", r)
	handler := cors.Default().Handler(r)
	http.Handle("/", handler)

	switch getEnv {
	case "development":
		fmt.Printf("main : Started : Listening on: http://localhost:8080")
		http.ListenAndServe("0.0.0.0:8080", nil)
	}
}
