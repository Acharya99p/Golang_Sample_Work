package main

import (
	"golang_rabbitmq/config"
	"fmt"
	"golang_rabbitmq/task_processor/app"
)

func main()  {
	fmt.Println("hello")
	config := config.GetConfig()
	fmt.Sprint()
	app := app.App{}
	app.Initialize(config)
	app.Run()
}
