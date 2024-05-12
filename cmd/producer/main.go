package main

import (
	"github.com/DiegoSenaa/go_rabbit/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "HELLO REI!", "amq.direct")
}
