package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "JLoka-log.log")
	fmt.Println("The Log-File is: ", LOGFILE)

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// jlokaLogger := log.New(f, "jloka", log.LstdFlags)
	jlokaLogger := log.New(f, "jloka ", log.Ldate|log.Ltime|log.Lshortfile)
	jlokaLogger.Println("Hi Jafar Loka Custom Logger World")
	jlokaLogger.Println("I am ITE Developer")
}