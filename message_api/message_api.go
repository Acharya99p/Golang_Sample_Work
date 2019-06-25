package message_api

type TaskDetails struct {
	Command string `json:"command"`
	Args []string `json:"args"`
}

type Message struct {
	TaskId uint `json:"message_id"`
	ReplyQueue string `json:"reply_queue"`
	Task TaskDetails `json:"task"`
}

type TaskResponse struct {
	TaskId uint `json:"task_id"`
	ExecStatus string `json:"exec_status"`
	ExecLogs string `json:"exec_logs"`
	ExecError string `json:"exec_error"`
	ExecTime float64 `json:"exec_time"`
	ExecHost string `json:"exec_host"`
}