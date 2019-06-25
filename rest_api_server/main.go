package main

import (
	"golang_rabbitmq/config"
	"golang_rabbitmq/rest_api_server/app"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
