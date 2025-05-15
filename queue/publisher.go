package queue

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func PublishResult(body []byte) {
	mqURL := os.Getenv("MQ_URL")
	if mqURL == "" {
		mqURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqp.Dial(mqURL)
	if err != nil {
		log.Printf("Failed to connect to MQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Channel error: %v", err)
		return
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		resultQueue, true, false, false, false, nil,
	)
	if err != nil {
		log.Printf("Queue declare error: %v", err)
		return
	}

	err = ch.Publish(
		"", resultQueue, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Publish error: %v", err)
	}
}
