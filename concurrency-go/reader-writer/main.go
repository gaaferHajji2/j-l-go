package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func writer(data *[]string, mutex *sync.Mutex) {
	for i := 0; ; i++ {
		mutex.Lock()
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

func reader(data *[]string, mutex *sync.Mutex, start time.Time) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		allData := getAllData(data)
		mutex.Unlock()
		fmt.Println("Data with length: ", len(allData), ", take: ", time.Since(start))
	}
}

func main() {
	data := make([]string, 0, 10000)
	mutex := sync.Mutex{}
	go writer(&data, &mutex)

	start := time.Now()
	for i := 0; i < 2000; i++ {
		go reader(&data, &mutex, start)
	}
	time.Sleep(50 * time.Second)
}
