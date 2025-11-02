package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("the id started is: %d at %s\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("the id ended is: %d at %s\n", id, time.Now().Format("15:04:05"))
}

func main() {
	for i := 0; i < 5; i++ {
		go doWork(i)
	}

	time.Sleep(2 * time.Second)
}
