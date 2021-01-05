package rabbit

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const (
	url = "amqp://guest:guest@localhost:5672/"
)

func Mains() {
	ch, q := createSender(nil, url)

	publish(nil, ch, q, "hello2")
	publish(nil, ch, q, "hello3")
	initRabbitReceive(nil, url)

}

func initRabbitReceive(logger *logrus.Logger, url string) {

	// connect to rabbit
	conn, err := amqp.Dial(url)
	failOnError(logger, err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// open a channel
	ch, err := conn.Channel()
	failOnError(logger, err, "Failed to open a channel")
	defer ch.Close()

	// declare the queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(logger, err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(logger, err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messageson queue: %s.", q.Name)
	<-forever

}

func failOnError(logger *logrus.Logger, err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
