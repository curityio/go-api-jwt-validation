package main

import (
	"encoding/json"
	"log"
	"os"
	"io/ioutil"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

var records *[]Record

var file = "./records.json"

// Load data from file
func readFile() {
	file, e := ioutil.ReadFile(file)
	if e != nil {
		log.Println("File error %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &records)
}

// Return all the records
func getRecords(w http.ResponseWriter, r *http.Request) {
	
	if(Authorize(w, r)){
		readFile()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	}
}

// Return record with given ID
func getRecord(w http.ResponseWriter, r *http.Request) {
	
	if(Authorize(w, r)){
		readFile()
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		for _, record := range *records {
			if strconv.Itoa(record.ID) == vars["id"] {
				json.NewEncoder(w).Encode(record)
				return
			}
		}

		json.NewEncoder(w).Encode(&Record{})
	}
}