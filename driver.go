package queueing

const (
	RABBITMQ_DRIVER = "rabbitmq"
)

type Driver interface {
	Connect(dsn string) error
	GetChannel(name string) (Channel, error)
}
