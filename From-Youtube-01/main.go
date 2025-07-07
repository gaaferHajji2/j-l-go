package main

import "fmt"

func main() {
	// fmt.Print("Hello Go World!\n")

	var task01 = "1. Watch the crash course of Go"

	var task02 = "2. Build network automation tools with Go"

	var task03 = "3. Build Microservices Applications With gRPC"

	var task04 = "4. Build my Applications using Go"

	task05 := "5. Test new type of declaring variables"

	// var taskSlice = []string{
	// 	"1. Watch the crash course of Go",
	// 	"2. Build network automation tools with Go",
	// 	"3. Build Microservices Applications With gRPC",
	// 	"4. Build my Applications using Go",
	// }

	var taskSlice = []string{task01, task02, task03, task04}

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

	fmt.Println("*** New Variable Declaration ***")
	fmt.Println(task05)

	fmt.Println("*** Task Of Items ***")
	fmt.Println("Tasks", taskSlice)
}
