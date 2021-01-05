package rabbit

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func failOnError2(logger *logrus.Logger, err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// url = "amqp://guest:guest@localhost:5672/"
func createSender(logger *logrus.Logger, url string) (*amqp.Channel, amqp.Queue) {

	// connect to rabbit
	conn, err := amqp.Dial(url)
	failOnError(logger, err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	// open a channel
	ch, err := conn.Channel()
	failOnError(logger, err, "Failed to open a channel")
	// defer ch.Close()

	// declare the queue.
	// Declaring a queue is idempotent - it will only be created if it doesn't exist already.
	fmt.Printf("declaring sender queue")
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(logger, err, "Failed to declare a queue")

	fmt.Printf("finished createSender \n")

	return ch, q
}

func publish(logger *logrus.Logger, ch *amqp.Channel, q amqp.Queue, msg string) {
	fmt.Printf("publishing %s\n", msg)

	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(logger, err, "Failed to publish a message")
}
