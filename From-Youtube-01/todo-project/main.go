package main

import "fmt"
import "net/http"

var task01 = "Watch the crash course of Go"
var task02 = "Build network automation tools with Go"
var task03 = "Build Microservices Applications With gRPC"
var task04 = "Build my Applications using Go"

var taskItems = []string{task01, task02, task03, task04}

func main() {
	// in this way we create handle for / without register it in ListenAndServe
	http.HandleFunc("/", helloLoka)
	http.HandleFunc("/show-tasks", showTasks);

	http.ListenAndServe("localhost:8080", nil);
}

func helloLoka(responseWriter http.ResponseWriter, request *http.Request) {
	var greeting = "Hi, My Name is Jafar Loka. Welcome to Todo project";
	fmt.Fprintln(responseWriter, greeting);
}

func showTasks(responseWriter http.ResponseWriter, request *http.Request) {

	printTasks(responseWriter, taskItems)

}

func printTasks(responseWriter http.ResponseWriter, taskItems []string) {

	fmt.Println("List of my Todos")

	for index, task := range taskItems {
		// fmt.Println(index+1, ".", task)
		fmt.Fprintf(responseWriter, "%d. %s\n", index+1, task)
	}
}
