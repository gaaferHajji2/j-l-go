package main

import "fmt"

func main() {
	// fmt.Print("Hello Go World!\n")

	var task01 = "1. Watch the crash course of Go"

	var task02 = "2. Build network automation tools with Go"

	var task03 = "3. Build Microservices Applications With gRPC"

	var task04 = "4. Build my Applications using Go"

	fmt.Println("##### Welcome to our Todolist App! #####")

	fmt.Println("*** List of my todos ***")

	fmt.Println(task01)

	fmt.Println(task02)

	fmt.Println(task03)

	fmt.Println(task04)

	fmt.Println()

	fmt.Println("*** List of my Tutorials ***")

	fmt.Println(task02)

	fmt.Println(task03)

	fmt.Println()

	fmt.Println("*** My Goals ***")
	fmt.Println(task04)
}
