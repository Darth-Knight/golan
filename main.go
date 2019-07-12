package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type bodyStruct struct {
	Name    string
	Address string
}

func saveData(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	decoder := json.NewDecoder(req.Body)
	var jsonData bodyStruct
	err := decoder.Decode(&jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error in json Parsing"))
	} else {
		response, err := writeDataToRedis(jsonData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(response))
		} else {
			response, err := pushDataInQueue(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(response))
			}
			w.Write([]byte(response))
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/saveData/", saveData).Methods("POST")
	log.Fatal(http.ListenAndServe(":1047", router))
}
