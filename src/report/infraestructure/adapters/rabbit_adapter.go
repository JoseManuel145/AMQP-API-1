package adapters

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitAdapter struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
}

// NewRabbitAdapter crea un nuevo adaptador conectándose a RabbitMQ, abriendo el canal y declarando la cola.
func NewRabbitAdapter(amqpURL, queueName string) (*RabbitAdapter, error) {
	conn, ch, err := ConnectRabbitMQ(amqpURL)
	if err != nil {
		return nil, err
	}

	// Declarar la cola
	_, err = ch.QueueDeclare(
		queueName, // nombre de la cola
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // argumentos
	)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitAdapter{
		Connection: conn,
		Channel:    ch,
		QueueName:  queueName,
	}, nil
}

// SendMessage publica un mensaje en la cola definida.
func (r *RabbitAdapter) SendMessage(mensaje string) error {
	err := r.Channel.Publish(
		"reports",    // exchange vacío usa el default
		"report_key", // routing key: el nombre de la cola
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(mensaje),
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Mensaje enviado a la cola %s: %s", r.QueueName, mensaje)
	return nil
}

// Close cierra el canal y la conexión.
func (r *RabbitAdapter) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Connection != nil {
		r.Connection.Close()
	}
}
