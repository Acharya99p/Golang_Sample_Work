package logger


import (
	log "github.com/Sirupsen/logrus"
	"os"
	"fmt"
)



func Initialize(logfile string){
	var filename string = logfile
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	Formatter :=  new(log.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	}else{
		log.SetOutput(f)
	}

	log.SetOutput(os.Stdout)
}


type GormLogger struct {}

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		log.WithFields(log.Fields{"module": "gorm", "type": "sql"}).Info(v[3])
	}
	if v[0] == "log" {
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Info(v[2])
	}
}


func Info(args ...interface{}){
	log.Info(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Warning(args ...interface{}) {
	log.Warning(args...)
}

