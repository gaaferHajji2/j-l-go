package main

import (
	"fmt"
	"runtime"
)

func t1() {
	fmt.Println("Hello from Jafar Loka")
}

func main() {
	go t1()
	go t1()
	go t1()
	// it is like yield, change the execution turn
	runtime.Gosched()
	fmt.Println("Hello Jafar Loka World")
}
