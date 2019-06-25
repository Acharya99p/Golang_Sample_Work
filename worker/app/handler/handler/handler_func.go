package handler

import (
	"os"
	"github.com/streadway/amqp"
	"golang_rabbitmq/task_processor/app/logger"
	"encoding/json"
	"golang_rabbitmq/utility_func"
	"golang_rabbitmq/message_api"
)

func GetHostname() string{

	hostname := os.Getenv("KUBENODE")
	if hostname != ""{
		return hostname
	}
	hostname, _ = os.Hostname()
	return hostname
}


func publish_message_to_create_env(ch *amqp.Channel, message message_api.TaskResponse, exchange string, route_key string)  {

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
