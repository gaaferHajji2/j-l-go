package main

import "fmt"

func main() {

	var task01 = "Watch the crash course of Go"

	var task02 = "Build network automation tools with Go"

	var task03 = "Build Microservices Applications With gRPC"

	var task04 = "Build my Applications using Go"

	var taskItems = []string{task01, task02, task03, task04}

	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("##### Welcome to our Todolist App! #####")

	for index, task := range taskItems {
		// fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
