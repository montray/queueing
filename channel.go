package queueing

const (
	TYPE_JSON = "application/json"
)

type Message struct {
	Body []byte
}

type Channel interface {
	Publish(contentType string, msg []byte) error
	Consume() (<-chan Message, error)
}