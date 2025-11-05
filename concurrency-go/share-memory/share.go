package main

import (
	"fmt"
	"time"
)

func t1(jloka *int) {
	for *jloka > 0 {
		time.Sleep(300 * time.Millisecond)
		*jloka -= 1
	}
}

func main() {
	jloka := 3
	start := time.Now()
	go t1(&jloka)
	for jloka > 0 {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("jloka is: ", jloka)
	}
	end := time.Now()
	fmt.Println("total is: ", end.Sub(start))
}
