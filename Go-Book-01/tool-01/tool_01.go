package main

import(
	"fmt"
	"os"
	"path/filepath"
	"os/exec"
)

func isExecutableViaLookPath(filename string) bool {
	_, err := exec.LookPath(filename)
	return err == nil
}


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

		// fmt.Println("The Fullpath is: ", fullpath)

		// Check if the file exists in that directory (full path)
		fileInfo, err := os.Stat(fullpath)

		if err != nil {
			continue
		} else {
			// fmt.Println("We don't have error: ")
		}

		mode := fileInfo.Mode()

		// Check if it's regular file (Not folder: useful for Linux-Systems)
		if !mode.IsRegular() {
			continue
		} else {
			// fmt.Println("Regular File")
		}

		// Check if the file is executable by os (For Linux Systems)
		if mode&0111 != 0 {
			fmt.Println("----------------------------")
			fmt.Println("The full path is: ", fullpath)
			fmt.Println("----------------------------")
			return
		} else {
			// fmt.Println("Not executable")
			// fmt.Println("The mode is: ", mode&0111)
		}
		
		// For Windows Systems
		if isExecutableViaLookPath(fullpath) {
			fmt.Println("----------------------------")
			fmt.Println("The full path is: ", fullpath)
			fmt.Println("----------------------------")
		}

		// fmt.Println("++++++++++++++++++++++++++++++++")
	}

	fmt.Println("The end of search")
}