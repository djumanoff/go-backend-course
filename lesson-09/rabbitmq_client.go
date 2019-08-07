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
	}

	sess := amqp.NewSession(cfg)
	err := sess.Connect()
	if err != nil {
		panic(err)
	}
	clt, err := sess.Client(amqp.ClientConfig{
		//RequestX:  "request",
		//ResponseX: "response",
	})
	if err != nil {
		panic(err)
	}

	resp, err := clt.Call("test", amqp.Message{Body: []byte("ping")})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Body))
}
