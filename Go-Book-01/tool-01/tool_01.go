package main

import(
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len (args) == 1 {
		fmt.Println("You must provide one argument")

		return
	}

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		
	}
}