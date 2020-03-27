package kafka

//go:generate mockgen -destination=kafkamock/kafkamock.go -package kafkamock . Client

type Client interface {
	WriteMessage(topic, message string) error
	ReadMessage(topic string) (string, error)
}
