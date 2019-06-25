package handler

import (
	"golang_rabbitmq/utility_func"
	"golang_rabbitmq/message_api"
	"encoding/json"
	"github.com/streadway/amqp"
	"golang_rabbitmq/task_processor/app/logger"
)


type Connection_details struct {
	Consumer_Details []ConsumerDetails `json:"consumer_details"`
}

type ArgumentsData struct {
	Hostname string `json:"hostname"`
	Identifier string `json:"identifier"`
}

type ConsumerDetails struct {
	Arguments ArgumentsData `json:"arguments"`
}


func publish_message_to_create_env(ch *amqp.Channel, message message_api.Message, exchange string, route_key string)  {

	logger.Info("sending message to",route_key)
	mess, _ := json.Marshal(&message)
	err := ch.Publish(
		exchange,         // exchange
		route_key, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        mess,
		})
	utility_func.FailOnError(err, "Failed to publish a message")
}
