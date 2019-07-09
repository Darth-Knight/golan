package main

import (
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func writeDataToRedis(t bodyStruct) (string, error) {
	log.Printf("Writing Data to redis started")
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Printf("Error while making connection with Redis", err)
		return "Error while making connection with redis", err
	}

	defer conn.Close()

	var key = generateRandomKey(20)

	var data = structToString(t)
	resp := conn.Cmd("SET", key, data)
	if resp.Err != nil {
		log.Printf("Error while Saving data in Redis", err)
		return "Error while Saving data in Redis", err
	}

	log.Printf("Writing Data to redis finished")
	return key, nil
}
