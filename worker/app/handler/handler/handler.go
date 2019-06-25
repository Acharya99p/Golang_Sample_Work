package handler

import (
	"github.com/streadway/amqp"
	"github.com/jinzhu/gorm"
	"github.com/michaelklishin/rabbit-hole"
	"golang_rabbitmq/task_processor/app/logger"
	"os/exec"
	"golang_rabbitmq/message_api"
	"encoding/json"
	"time"
)

func Env_create_parser(db *gorm.DB, ch *amqp.Channel, rmqc *rabbithole.Client, message *amqp.Delivery)  {
	logger.Info("task received", string(message.Body))


	var mess message_api.Message

	json.Unmarshal(message.Body, &mess)


	logger.Info("start executing", mess)
	start := time.Now()
	cmd := exec.Command(mess.Task.Command, mess.Task.Args...)
	out, err  := cmd.CombinedOutput()
	var interval float64
	interval = time.Since(start).Seconds()

	exec_status := "success"
	exec_error := ""
	if err != nil{
		exec_status = "failed"
		exec_error = err.Error()
	}
	logger.Info(exec_status)
	reply_mess := message_api.TaskResponse{TaskId:mess.TaskId, ExecHost:GetHostname(), ExecStatus:exec_status,
	              ExecError:exec_error, ExecLogs:string(out), ExecTime:interval}

	logger.Info("Output",reply_mess)

	publish_message_to_create_env(ch, reply_mess, "tasktopic", mess.ReplyQueue )

	message.Ack(false)
}
