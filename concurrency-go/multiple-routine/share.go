package main

import (
	"fmt"
	"time"
)

func spend(t1 *int) {
	for i := 0; i < 1000000; i++ {
		*t1 -= 10
	}

	fmt.Println("Spend completed")
}

func save(t1 *int) {
	for i := 0; i < 1000000; i++ {
		*t1 += 10
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
