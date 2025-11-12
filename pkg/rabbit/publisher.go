package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel *amqp.Channel
}

func NewPublisher(url string) (*Publisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}

	return &Publisher{channel: ch}, nil
}

func (p *Publisher) Publish(queueName string, body []byte) {
	_, err := p.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	err = p.channel.Publish(
		"", queueName, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Printf("failed to publish message: %v", err)
	}
}
