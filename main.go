package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":3000"
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/saveData/", saveData).Methods("POST")
	port := getPort()
	log.Fatal(http.ListenAndServe(port, router))
}
