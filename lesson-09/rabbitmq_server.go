package main

import (
	"fmt"

	"github.com/djumanoff/amqp"
)

func main() {
	cfg := amqp.Config{
		Host:        "localhost",
		Port:        5672,
		VirtualHost: "/",
		User:        "guest",
		Password:    "guest",
		LogLevel:    5,
	}

	sess := amqp.NewSession(cfg)
	err := sess.Connect()
	if err != nil {
		panic(err)
	}
	server, err := sess.Server(amqp.ServerConfig{
		//RequestX:  "request",
		//ResponseX: "response",
	})
	if err != nil {
		panic(err)
	}

	server.Endpoint("test", func(message amqp.Message) *amqp.Message {
		fmt.Println(string(message.Body))
		return &amqp.Message{Body: []byte("response")}
	})

	server.Start()
}
