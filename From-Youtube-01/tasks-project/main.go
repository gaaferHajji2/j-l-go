package main

import "fmt"

func main() {

	fmt.Println("##### Welcome to our Todolist App! #####")

	var task01 = "Watch the crash course of Go"
	var task02 = "Build network automation tools with Go"
	var task03 = "Build Microservices Applications With gRPC"
	var task04 = "Build my Applications using Go"

	var taskItems = []string{task01, task02, task03, task04}

	printTasks(taskItems)
	fmt.Println()
	fmt.Println("### Updated List ###")
	taskItems = addTask(taskItems, "Build network applications")
	printTasks(taskItems)
}

func printTasks(taskItems []string) {

	fmt.Println("List of my Todos")

	for index, task := range taskItems {
		// fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)

	return updatedTaskItems
}
