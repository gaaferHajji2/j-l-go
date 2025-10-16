package main

import (
	"fmt"
	s "strings"
	// "unicode"
)

func main() {
	var f = fmt.Printf

	f("To Upper: %s\n", s.ToUpper("Jafar Loka is ITE Develoeper"))
	f("To Lower: %s\n", s.ToLower("Jafar Loka is ITE Develoeper"))
	f("Title: %s\n", s.Title("Jafar Loka is ITE Develoeper"))
	f("EqualFold: %v\n", s.EqualFold("JLoka", "jloka"))
	f("Index: %v\n", s.Index("Jafar Loka is ITE Develoeper", "Jafar Loka"))
}