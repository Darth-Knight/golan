package main

import (
	"log"

	"github.com/streadway/amqp"
)

func pushDataInQueue(key string) (string, error) {
	log.Println("Pushing key -- ", key, "in Queue")

	conn, err := amqp.Dial("amqp://ykbsznqg:KSZVK1JlGtYgPmmzDI7I9rZfiB1WUEPY@crane.rmq.cloudamqp.com/ykbsznqg")
	if err != nil {
		log.Println("Failed to connect to RabbitMQ", err)
		return "Failed to connect to RabbitMQ", err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel in queue", err)
		return "Failed to open a channel in queue", err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"redisKeys", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	body := key
	err = ch.Publish(
		"", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("Sent %s", body)
	if err != nil {
		log.Println("Failed to push key in queue", err)
		return "Failed to push key in queue", err
	}
	var response = "Key- " + key + " entered in Queue"
	log.Printf(response)
	return response, nil

}
