package main

import (
	"log/syslog"
	"log"
)

func main() {
	jlokaLog, err := syslog.New(syslog.LOG_WARNING|syslog.LOG_DAEMON, "my-app")

	if err != nil {
		log.Println("The err is: ", err)
		return
	}

	defer jlokaLog.Close()

	log.SetOutput(jlokaLog)
	log.Println("JLoka Logging-01")
}