package main

import (
	"fmt"
	"sync"
	"time"
)

func spend(t1 *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*t1 -= 10
		mutex.Unlock()
		// From Doc: Gosched yields the processor, allowing other goroutines to run.
		// It does not suspend the
		// current goroutine, so execution resumes automatically.
		// runtime.Gosched()
	}

	fmt.Println("Spend completed")
}
func save(t1 *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*t1 += 10
		mutex.Unlock()
		// here the goroutine runs on multiple processors, so the result also has race conditions
		// runtime.Gosched()
	}
	fmt.Println("Save completed")
}
func main() {
	t1 := 500
	// When we create a new mutex, its initial state is always unlocked.
	mutex := sync.Mutex{}
	go save(&t1, &mutex)
	go spend(&t1, &mutex)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Println("Total is: ", t1)
	mutex.Unlock()
}
