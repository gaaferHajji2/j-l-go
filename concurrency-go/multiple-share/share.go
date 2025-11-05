package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const characters = "abcdefghijklmnopqrstuvwxyz"

func CheckLetters(url string, freq []int) {
	resp, err := http.Get(url)
	if err != nil {
		panic("Error in Get")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Error with status: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	for _, c := range body {
		t1 := strings.ToLower(string(c))
		t1Index := strings.Index(characters, t1)
		if t1Index >= 0 {
			freq[t1Index] += 1
		}
	}

	fmt.Println(url, " completed")
}

func main() {
	freq := make([]int, 26)
	url := "https://example.com"
	for i := 0; i < 10; i++ {
		CheckLetters(url, freq)
	}

	for t1, t2 := range characters {
		fmt.Printf("%c is: %d\n", t2, freq[t1])
	}
}
