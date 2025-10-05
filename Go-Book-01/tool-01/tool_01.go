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
		fullpath := filepath.Join(directory, file)

		// Check if the file exists in that directory (full path)
		fileInfo, err := os.Stat(fullpath)

		if err != nil {
			continue
		}

		mode := fileInfo.Mode()

		// Check if it's regular file (Not folder: useful for Linux-Systems)
		if !mode.IsRegular() {
			continue
		}

		// Check if the file is executable by os
		if mode&0111 != 0 {
			fmt.Println("The full path is: ", fullpath)
			break
		}
	}
}