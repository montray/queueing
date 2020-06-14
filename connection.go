package queueing

type Connection struct {
	driver Driver
	dsn string
}

func NewConnection(config Config) (*Connection, error) {
	switch config.Driver {
	case RABBITMQ_DRIVER:
		conn := &Connection{
			driver: &RabbitMqDriver{},
			dsn: config.Dsn,
		}

		err := conn.dial()

		if err != nil {
			return nil, err
		}

		return conn, nil
	default:
		panic("not known driver")
	}
}

func (conn *Connection) dial() error {
	err := conn.driver.Connect(conn.dsn)
	return err
}

func (conn *Connection) GetChannel(channel string) (Channel, error) {
	return conn.driver.GetChannel(channel)
}