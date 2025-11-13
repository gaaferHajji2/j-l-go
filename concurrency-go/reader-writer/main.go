package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func writer(data *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock() // this will lock for wrting only
		*data = append(*data, "Data with index: "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Data has been added")
	}
}

func getAllData(data *[]string) []string {
	copyData := make([]string, 0, len(*data))

	for _, d := range *data {
		copyData = append(copyData, d)
	}
	return copyData
}

func reader(data *[]string, mutex *sync.RWMutex, start time.Time) {
	for i := 0; i < 1000; i++ {
		mutex.RLock()
		allData := getAllData(data)
		mutex.RUnlock()
		runtime.Gosched()
		fmt.Println("Data with length: ", len(allData), ", take: ", time.Since(start))
	}
}

func main() {
	data := make([]string, 0, 10000)
	mutex := sync.RWMutex{}
	go writer(&data, &mutex)

	start := time.Now()
	for i := 0; i < 2000; i++ {
		go reader(&data, &mutex, start)
	}
	time.Sleep(50 * time.Second)
}
