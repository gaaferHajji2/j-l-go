package main

import (
	"log"
	"os"
	"io"
)

func SetupSimpleWriter(logFile string) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile )
}

func main() {
	SetupSimpleWriter("jloka.log")
	log.Println("JLoka Simple Logger For Windows")
	log.Printf("Hello %s Logging World", "Jafar Loka")
}