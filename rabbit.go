package queueing

import "github.com/streadway/amqp"

type RabbitMqMessage struct {
	delivery amqp.Delivery
	body []byte
}

func (r RabbitMqMessage) GetBody() []byte {
	return r.body
}

func (r RabbitMqMessage) Ack() error {
	return r.delivery.Ack(false)
}

type RabbitMqChannel struct {
	name string
	ch *amqp.Channel
}

func (r *RabbitMqChannel) Publish(contentType string, msg []byte) error {
	err := r.ch.Publish(
		"",
		r.name,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body: msg,
		},
	)

	return err
}

func (r *RabbitMqChannel) Consume() (<-chan Message, error) {
	msgs, err := r.ch.Consume(r.name, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	out := make(chan Message)

	go func(in <-chan amqp.Delivery, out chan Message) {
		for {
			del := <-in
			msg := RabbitMqMessage{
				delivery: del,
				body: del.Body,
			}
			out <- msg
		}
	}(msgs, out)

	return out, err
}

type RabbitMqDriver struct {
	amqp *amqp.Connection
}

func (r *RabbitMqDriver) Connect(dsn string) error {
	var err error
	r.amqp, err = amqp.Dial(dsn)
	return err
}

func (r *RabbitMqDriver) GetChannel(name string) (Channel, error) {
	ch, err := r.amqp.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &RabbitMqChannel{ch: ch, name: name}, nil
}