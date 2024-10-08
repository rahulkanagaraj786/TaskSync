package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Seperate Routine to send MQ message
func sendToMQ() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()
	q, _ := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	// count := 0
	for {
		message := <-msgCh
		// count++
		// fmt.Println("count : ", count, " Time:", currentTime())
		fmt.Println("Sending message:", message)
		ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})
	}
}