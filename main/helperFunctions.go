package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomKey(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func structToString(t bodyStruct) string {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(b))
	return string(b)
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
