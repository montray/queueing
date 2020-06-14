package queueing

const (
	TYPE_JSON = "application/json"
)

type ChannelConfig struct {
	AutoAck bool
}

type Message interface {
	GetBody() []byte
	Ack() error
}

type Channel interface {
	Publish(contentType string, msg []byte) error
	Consume() (<-chan Message, error)
}