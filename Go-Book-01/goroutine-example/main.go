package main

import (
	"fmt"
	"time"
)

func printFunc(name string, index int) {
	fmt.Printf("My Name is: %s-%d\n", name, index)
}

func main() {
	for i:=0; i<3; i++ {
		go printFunc("Jafar Loka", i)
	}

	time.Sleep(2 * time.Second)
}