package main

import (
    "encoding/json"
    "github.com/masnun/gopher-and-rabbit"
    "github.com/streadway/amqp"
    "log"
    "math/rand"
    "time"
)

func handleError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }

}

func main() {
    conn, err := amqp.Dial(gopher_and_rabbit.Config.AMQPConnectionURL)
    handleError(err, "Can't connect to AMQP")
    defer conn.Close()

    amqpChannel, err := conn.Channel()
    handleError(err, "Can't create a amqpChannel")

    defer amqpChannel.Close()

    queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
    handleError(err, "Could not declare `add` queue")

    rand.Seed(time.Now().UnixNano())

    addTask := gopher_and_rabbit.AddTask{Number1: rand.Intn(999), Number2: rand.Intn(999)}
    body, err := json.Marshal(addTask)
    if err != nil {
        handleError(err, "Error encoding JSON")
    }

    err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
        DeliveryMode: amqp.Persistent,
        ContentType:  "text/plain",
        Body:         body,
    })

    if err != nil {
        log.Fatalf("Error publishing message: %s", err)
    }

    log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)

}



func main() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "hello", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    body := "Hello World!"
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    log.Printf(" [x] Sent %s", body)
    failOnError(err, "Failed to publish a message")
}




 //    msgs, err := ch.Consume(
    //  q.Name, // queue
    //  "",     // consumer
    //  true,   // auto-ack
    //  false,  // exclusive
    //  false,  // no-local
    //  false,  // no-wait
    //  nil,    // args
    // )
    // handleError(err, "Failed to register a consumer")

    // forever := make(chan bool)

    // go func() {
    //  for d := range msgs {
    //      log.Println("Received a message: %s", d.Body)
    //  }
    // }()

    // log.Println(" [*] Waiting for messages. To exit press CTRL+C")
    // <-forever