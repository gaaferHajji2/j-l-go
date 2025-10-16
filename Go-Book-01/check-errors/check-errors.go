package main

import (
	"errors"
	"os"
	"strconv"
	"fmt"
)

func CheckError(a int, b int) error {
	if a == 0 && b == 0 {
		return errors.New("JLoka Custom Error Msg")
	}
	return nil
}

func CheckErrorWithFormat(a int, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d && b %d && user uid %d", a, b, os.Getuid())
	}
	return nil
}

func main() {
	err := CheckError(0, 0)
	if err != nil {
		fmt.Println("The error is: ", err)
	}

	// here we reuse err, so we set = instead of :=
	err = CheckErrorWithFormat(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	_, err = strconv.Atoi("J-01")
	if err != nil {
		fmt.Println("the err is: ", err)
	}
}