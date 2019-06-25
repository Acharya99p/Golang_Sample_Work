package message_handler

import (
	"github.com/streadway/amqp"
	"golang_rabbitmq/utility_func"
	"fmt"
)

var Ch *amqp.Channel

var Conn *amqp.Connection

func Initialize()  {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	utility_func.FailOnError(err, "Failed to connect to RabbitMQ")
	Ch, err = Conn.Channel()
	utility_func.FailOnError(err, "Failed to open a channel")

}

func Send_env_message(id uint)  {

	fmt.Println("sending message with ", fmt.Sprint(id))

	if Ch != nil {
		Ch.Publish(
			"tasktopic",      // exchange
			"TaskQueue", // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         [] byte(fmt.Sprint(id)),
			})
	}
}