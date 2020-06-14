package main

import (
	"fmt"
	"github.com/montray/queueing"
	"log"
)

func main() {
	q, err := queueing.NewConnection(queueing.Config{
		Driver: queueing.RABBITMQ_DRIVER,
		Dsn: "amqp://user:user@localhost:5672/",
	})

	if err != nil {
		log.Fatal(err)
	}

	ch, err := q.GetChannel("example_ch")

	if err != nil {
		log.Fatal(err)
	}

	err = ch.Publish(queueing.TYPE_JSON, []byte("example message"))
	err = ch.Publish(queueing.TYPE_JSON, []byte("example message_2"))

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume()

	if err != nil {
		log.Fatal(err)
	}

	for v := range msgs {
		fmt.Println(string(v.Body))
	}
}
