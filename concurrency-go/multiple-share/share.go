package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const characters = "abcdefghijklmnopqrstuvwxyz"

func CheckLetters(url string, freq []int, mutex *sync.Mutex) {
	resp, err := http.Get(url)
	if err != nil {
		panic("Error in Get")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Error with status: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	mutex.Lock()
	for _, c := range body {
		t1 := strings.ToLower(string(c))
		t1Index := strings.Index(characters, t1)
		if t1Index >= 0 {
			freq[t1Index] += 1
		}
	}
	mutex.Unlock()
	fmt.Println(url, " completed")
}

func main() {
	freq := make([]int, 26)
	url := "https://jsonplaceholder.typicode.com/users/"
	mutex := sync.Mutex{}
	for i := 1; i < 11; i++ {
		// this will may cause race conditions
		go CheckLetters(url+strconv.Itoa(i), freq, &mutex)
	}
	time.Sleep(5 * time.Second)
	mutex.Lock()
	for t1, t2 := range characters {
		fmt.Printf("%c is: %d\n", t2, freq[t1])
	}
	mutex.Unlock()
}
