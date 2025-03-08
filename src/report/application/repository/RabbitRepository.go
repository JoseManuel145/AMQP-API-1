package rabbit

import "report/src/report/infraestructure/adapters"

// IRabbit define el contrato para enviar mensajes.
type IRabbit interface {
	SendMessage(mensaje string) error
}

// RabbitRepository es la implementación de IRabbit utilizando el RabbitAdapter.
type RabbitRepository struct {
	adapter *adapters.RabbitAdapter
}

// NewRabbitRepository crea una nueva instancia de RabbitRepository a partir de un adaptador.
func NewRabbitRepository(adapter *adapters.RabbitAdapter) *RabbitRepository {
	return &RabbitRepository{
		adapter: adapter,
	}
}

// SendMessage delega el envío de mensaje al adaptador.
func (r *RabbitRepository) SendMessage(mensaje string) error {
	return r.adapter.SendMessage(mensaje)
}
