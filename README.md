# Golang_Sample_Work

Remote Executer

This repo is a mono repo of multiple components of remote executer.

Introduction:
If you have thousands of servers in your infra and you want to run some jobs on all or some of them for computing or
infra management. This project utilize rabbitmq topic exchange.


Architecture:

                +----------------------+
     +----------+   task_processor     |
     |          +-----+-------+--------+
     |                ^       |
     |                |       |                worker     worker   worker
+----+--+             |       |                   +         +        +
|   DB  |             |       |                   |         |        |
+-+-----+         +---+-------v-------------------+---------+--------+---+
  |               |                  RAbbitMQ                            |
  |               +-^-----------------------------+---------+--------+---+
  |                 |                             |         |        |
  |                 |                             |         |        |
  |     +-----------+-----+                       +         +        +
  +-----+ rest_api_server |                     worker    worker   worker
        +-----------------+


Setup:
First set up your golang https://golang.org/doc/install

cd <project dir>

go get ./...

Please set your configs for the project in config/

pushd rest_api_server
    go build main.go
    ./main
popd

pushd task_processor
    go build main.go
    ./main
popd

pushd worker
    go build main.go
    ./main
popd


You can run worker on any machine, just make sure machine can reach to rabbitmq server



Api available

/tasks GET
/tasks POST
/tasks/<id> GET
/tasks/<id> DELETE
/tasks/<id> UPDATE


task json for POST
{
	"title" : "sleep",
	"host" : "common",
	"details" : {
		"command" : "/bin/sleep",
		"args" : ["10"]
	}
}

