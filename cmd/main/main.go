package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/TheRAWagent/go-bookstore/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("PORT not set in .env file")
	}
	log.Fatal(http.ListenAndServe(":" + port, router))
}