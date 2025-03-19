package core

import (
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// Ejemplo de como se implementaria la conexion a RabbitMQ con multiples entidades
func NewRabbitMQClient(amqpURL string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQClient{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func (r *RabbitMQClient) Close() {
	r.Channel.Close()
	r.Conn.Close()
}
