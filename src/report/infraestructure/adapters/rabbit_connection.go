package adapters

import (
	"log"

	"github.com/streadway/amqp"
)

// ConnectRabbitMQ establece la conexión y el canal con RabbitMQ.
func ConnectRabbitMQ(amqpURL string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	log.Println("Conexión a RabbitMQ establecida.")
	return conn, ch, nil
}
