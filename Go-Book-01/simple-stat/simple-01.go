package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
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

	fmt.Printf("The number of values is: #%d\n", nValues)
	fmt.Println("The minimum value is: ", min)
	fmt.Println("The maximum value is: ", max)
	fmt.Println("The sum of all numbers: ", sum)

	if nValues == 0 {
		return
	}

	meanValue := sum / float64(nValues)
	fmt.Printf("The mean value is: %.5f\n", meanValue)

	var squared float64
	for i := 1; i< len(args); i++ {
		t3, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}

		squared = squared + math.Pow((t3-meanValue), 2)
	}

	standardDevision := math.Sqrt(squared / float64(nValues))
	fmt.Printf("The std devision is: %.5f", standardDevision)
}