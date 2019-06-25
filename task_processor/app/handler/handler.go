package handler

import (
	"github.com/streadway/amqp"
	"golang_rabbitmq/model"
	"github.com/jinzhu/gorm"
	"github.com/michaelklishin/rabbit-hole"
	"golang_rabbitmq/message_api"
	"encoding/json"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"golang_rabbitmq/task_processor/app/logger"
)

func Env_create_parser(db *gorm.DB, ch *amqp.Channel, rmqc *rabbithole.Client, message *amqp.Delivery)  {
	logger.Info("Got Task id for processing",string(message.Body))
	id := string(message.Body)
	var task  model.Task
	time.Sleep(50 * time.Millisecond)
	if err := db.First(&task, id).Error; err != nil {
		logger.Warning("Tasks not fount", err.Error())
	}
	logger.Info("job to be sent is ", task)
	var taskdetails message_api.TaskDetails
	json.Unmarshal(task.Details.RawMessage, &taskdetails)
	mess := message_api.Message{
		TaskId: task.ID,
		ReplyQueue: "TaskResponseQueue",
        Task: taskdetails,
	}
	if task.Host != "common"{
		task.Host = task.Host
	}
	publish_message_to_create_env(ch, mess,"tasktopic", task.Host)
	message.Ack(false)
}

func Reply_queue_parser(db *gorm.DB, ch *amqp.Channel, rmqc *rabbithole.Client, message *amqp.Delivery)  {
	logger.Info("Got Reply messsage from agent",string(message.Body))
	reply_mess := message_api.TaskResponse{}
	json.Unmarshal([]byte(message.Body), &reply_mess)

		var env_job_host model.TaskHost
		db.Where(model.TaskHost{TaskId:reply_mess.TaskId, Host: reply_mess.ExecHost}).FirstOrCreate(&env_job_host)
		details, _ := json.Marshal(reply_mess)
		env_job_host.Response = postgres.Jsonb{details}
		db.Save(env_job_host)

	message.Ack(false)
}

