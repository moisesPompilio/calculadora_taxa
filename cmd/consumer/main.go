package main

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/moisesPompilio/calculadora_taxa/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)
	for msg := range out {
		println(string(msg.Body))
	}
}
