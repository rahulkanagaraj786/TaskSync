package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Separate goroutine for sending messages to RabbitMQ
func mqMessageSender() {
	// Establish connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Open a channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("[ERROR] Failed to open a channel: %v", err)
	}
	defer channel.Close()

	// Declare a queue
	queue, err := channel.QueueDeclare(
		"taskQueue", // Queue name
		false,       // Durable
		false,       // Auto-delete when unused
		false,       // Exclusive
		false,       // No-wait
		nil,         // Additional arguments
	)
	if err != nil {
		log.Fatalf("[ERROR] Failed to declare a queue: %v", err)
	}

	// Infinite loop for sending messages from the channel
	for {
		message := <-jobQueue // Updated channel name
		fmt.Println("Sending job message:", message)

		// Publish the message to RabbitMQ
		err = channel.Publish(
			"",         // Default exchange
			queue.Name, // Routing key (queue name)
			false,      // Mandatory flag
			false,      // Immediate flag
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err != nil {
			log.Printf("[ERROR] Failed to publish message: %v", err)
		}
	}
}
