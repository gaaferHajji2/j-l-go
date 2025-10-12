package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("At least one arg for program")
		return
	}

	var min, max float64
	var initialized = true

	nValues := 0

	var sum float64

	for i := 0; i<len(args); i++ {
		t1, err := strconv.ParseFloat(args[i], 64)

		if err!= nil {
			continue
		}

		nValues = nValues + 1;
		sum = sum + t1

		if initialized {
			min = t1
			max = t1
			initialized = false
			continue
		}

		if t1 < min {
			min = t1
		}

		if t1 > max {
			max = t1
		}
	}

	fmt.Println("The number of values is: #", nValues)
	fmt.Println("The minimum value is: ", min)
	fmt.Println("The maximum value is: ", max)
	fmt.Println("The sum of all numbers: ", sum)
}