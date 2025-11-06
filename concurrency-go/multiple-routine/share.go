package main

import (
	"fmt"
	"runtime"
	"time"
)

func spend(t1 *int) {
	for i := 0; i < 1000000; i++ {
		*t1 -= 10
		runtime.Gosched()
	}

	fmt.Println("Spend completed")
}
func save(t1 *int) {
	for i := 0; i < 1000000; i++ {
		*t1 += 10
		// here the goroutine runs on multiple processors, so the result also has race conditions
		runtime.Gosched()
	}
	fmt.Println("Save completed")
}
func main() {
	t1 := 500
	go save(&t1)
	go spend(&t1)
	time.Sleep(2 * time.Second)
	fmt.Println("Total is: ", t1)
}
