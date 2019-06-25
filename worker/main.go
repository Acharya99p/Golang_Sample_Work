package main

import (
	"golang_rabbitmq/config"
	"fmt"
	"golang_rabbitmq/worker/app"
)

func main()  {
	config := config.GetConfig()
	fmt.Sprint()
	app := app.App{}
	app.Initialize(config)
	app.Run()
}
