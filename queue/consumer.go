package queue

import (
	"log"

	"github.com/streadway/amqp"
)

var resultQueue = "ocr_results"

func Consume(mqURL string, handler func([]byte)) error {
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"ocr_tasks", true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		"ocr_tasks", "", true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()

	log.Println("Listening for OCR tasks...")
	select {} // блокируем завершение main
}
