package repositories

type IRabbit interface {
	SendMessage(mensaje string) error
}
