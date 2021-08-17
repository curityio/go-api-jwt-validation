package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	//Load config file
	err := godotenv.Load()
	
	//Create Algorithm interface to use for JWT verification
	setAlgorithm(os.Getenv("JWKS"))

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/api/records", getRecords)
	router.HandleFunc("/api/records/{id}", getRecord)

	port := os.Getenv("PORT")
	log.Println("Serving on port:", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}