package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
	// "github.com/mediocregopher/radix.v2/redis"
)

func writeDataToRedis(t bodyStruct) (string, error) {
	log.Printf("Writing Data to redis started")
	conn, err := redis.DialURL("redis://h:p7691e37f22f3598216226f0a3e50eef070847ba525f6d9f55415279f062c127c@ec2-52-202-172-13.compute-1.amazonaws.com:25289")
	if err != nil {
		log.Printf("Error while making connection with Redis", err)
		return "Error while making connection with redis", err
	}

	defer conn.Close()

	var key = generateRandomKey(20)

	var data = structToString(t)
	_, err2 := conn.Do("SET", key, data)
	if err2 != nil {
		log.Printf("Error while Saving data in Redis", err2)
		return "Error while Saving data in Redis", err2
	}

	log.Printf("Writing Data to redis finished")
	return key, nil
}
