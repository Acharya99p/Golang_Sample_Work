package app

import (
	"golang_rabbitmq/config"
	"github.com/streadway/amqp"
	"fmt"
	"golang_rabbitmq/utility_func"
	"github.com/jinzhu/gorm"
	"log"
	_ "golang_rabbitmq/model"
	"github.com/michaelklishin/rabbit-hole"
	"golang_rabbitmq/task_processor/app/handler"
	"golang_rabbitmq/task_processor/app/logger"
)

type App struct {
	Ch *amqp.Channel
	Conn *amqp.Connection
	DB     *gorm.DB
	Rmqc *rabbithole.Client
}

func (a *App)Initialize(config *config.Config)  {
	logger.Initialize("logfile.log")
	Uri := fmt.Sprintf("%s://%s:%s@%s:%s",
		config.RmQ.Dialect,
		config.RmQ.Username,
		config.RmQ.Password,
		config.RmQ.Host,
		config.RmQ.Port,
	)
	var err error
	a.Conn, err = amqp.Dial(Uri)
	utility_func.FailOnError(err, "Failed to connect to RabbitMQ")
	a.Ch, err = a.Conn.Channel()
	utility_func.FailOnError(err, "Failed to open a channel")
	a.DB, err = gorm.Open(config.DB.Dialect, config.DB.Name)
	if err != nil {
		log.Fatal("Could not connect database")
	}
	a.DB.SetLogger(&logger.GormLogger{})
	//a.DB.LogMode(true)
	a.Rmqc, err = rabbithole.NewClient("http://"+ config.RmQ.Host + ":" + config.RmQ.ApiPort, config.RmQ.Username, config.RmQ.Password)
	utility_func.FailOnError(err, "Unable to create Rmqc")
}

func (a *App)declarequeue(queue_name string)  {
	_ , err := a.Ch.QueueDeclare(
		queue_name,    // name
		true,       // durable
		true,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utility_func.FailOnError(err, "Failed to declare a queue")

}


func (a *App) bind(exchange_name string, queue_name string, bind_route_key string, f func(db *gorm.DB, ch *amqp.Channel, rmqc *rabbithole.Client,message *amqp.Delivery))  {

	err := a.Ch.QueueBind(
		queue_name,        // queue name
		bind_route_key,             // routing key
		exchange_name, // exchange
		false,
		nil)
	utility_func.FailOnError(err, "Failed to bind a queue")



	msgs, err := a.Ch.Consume(
		queue_name, // queue
		"",     // consumer
		false,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	a.bind_handlers(msgs, f)
}


func (a *App)exchange_init(exchange_name string) {
	err := a.Ch.ExchangeDeclare(
		exchange_name, // name
		"topic",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	utility_func.FailOnError(err, "Failed to declare an exchange")

	err = a.Ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	utility_func.FailOnError(err, "Failed to set QoS")

}

func (a *App)bind_handlers(msgs<-chan amqp.Delivery, f func(db *gorm.DB, ch *amqp.Channel, rmqc *rabbithole.Client,message *amqp.Delivery))  {
	go func() {
		for d := range msgs {
			f(a.DB,a.Ch,a.Rmqc,&d)
		}
	}()
}



func (a *App)Run()  {
	ExchangeName := "tasktopic"
	a.exchange_init(ExchangeName)
	a.declarequeue("TaskQueue")
	a.bind(ExchangeName, "TaskQueue","TaskQueue", handler.Env_create_parser)
	a.declarequeue("TaskResponseQueue")
	a.bind(ExchangeName, "TaskResponseQueue","TaskResponseQueue", handler.Reply_queue_parser)
	//a.declarequeue("new_node")
	//a.bind(ExchangeName,"new_node","new_node",handler.New_node_message)
//	a.declarequeue("all_listener")
//	a.bind(ExchangeName,"all_listener","#",handler.All_Messages)
	//a.bind(ExchangeName,"TaskQueue","#",handler.All_Messages)
	forever := make(chan bool)
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}