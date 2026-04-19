package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// =====================================================
// 3. STRUCT DEFINITION
// =====================================================
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// =====================================================
// 6. GENERICS - Custom numeric constraint
// =====================================================
type Numeric interface {
	int | float64
}

func Sum[T Numeric](numbers ...T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

// =====================================================
// 6. POLYMORPHISM - Interface + multiple implementations
// =====================================================
type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hi, I'm " + p.Name + " (a person)"
}

type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "Beep boop! I am " + r.Model + " (a robot)"
}

// =====================================================
// 7. CONTROL STRUCTURES DEMO FUNCTION
// =====================================================
func demonstrateControlStructures() {
	fmt.Println("\n=== 7. CONTROL STRUCTURES DEMO ===")

	// if / else if / else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// for loop (classic C-style)
	fmt.Print("For loop (0 to 4): ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// for range over slice
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("For range over slice:")
	for index, fruit := range fruits {
		fmt.Printf("  Index %d: %s\n", index, fruit)
	}

	// for range over map
	colors := map[string]string{"red": "#FF0000", "green": "#00FF00"}
	fmt.Println("For range over map:")
	for key, value := range colors {
		fmt.Printf("  %s -> %s\n", key, value)
	}

	// switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// switch with initialization and expression
	switch hour := 14; {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func main() {
	fmt.Println("🚀 GO LANG CORE FEATURES COMPLETE DEMO")
	fmt.Println("=====================================")

	// =====================================================
	// 2. INTEGERS, FLOATS, STRINGS, CHARS (RUNES)
	// =====================================================
	fmt.Println("\n=== 2. BASIC DATA TYPES ===")
	var integer int = 42
	var float32Val float32 = 3.14159
	var float64Val float64 = 2.718281828459045 // double precision
	var text string = "Hello, GoLang! 🌟"
	var character rune = '🔥' // Go's "char" is rune (Unicode code point)

	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float32: %.5f\n", float32Val)
	fmt.Printf("Float64 (double): %.15f\n", float64Val)
	fmt.Printf("String: %s\n", text)
	fmt.Printf("Rune/Char: %c (Unicode: %U)\n", character, character)

	// =====================================================
	// 3. ARRAYS, SLICES, STRUCTS + INDEX ACCESS
	// =====================================================
	fmt.Println("\n=== 3. ARRAYS, SLICES, STRUCTS ===")

	// Fixed-size array
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array (fixed size): %v\n", array)
	fmt.Printf("Array element at index 2: %d\n", array[2])
	fmt.Printf("Array length: %d\n", len(array))

	// Dynamic slice
	slice := []string{"Go", "is", "fast", "and", "simple"}
	slice = append(slice, "and concurrent!")
	fmt.Printf("Slice (dynamic): %v\n", slice)
	fmt.Printf("Slice element at index 0: %s\n", slice[0])
	fmt.Printf("Slice length: %d, Capacity: %d\n", len(slice), cap(slice))

	// Struct
	user := User{
		ID:   1,
		Name: "Alice",
		Age:  28,
	}
	fmt.Printf("Struct: %+v\n", user)
	fmt.Printf("Struct field access: Name = %s, Age = %d\n", user.Name, user.Age)

	// =====================================================
	// 4. WRITE JSON OBJECT TO FILE
	// =====================================================
	fmt.Println("\n=== 4. WRITING JSON TO FILE ===")

	users := []User{
		{ID: 1, Name: "Alice", Age: 28},
		{ID: 2, Name: "Bob", Age: 35},
		{ID: 3, Name: "Charlie", Age: 22},
	}

	// Create file
	file, err := os.Create("users.json")
	if err != nil {
		fmt.Printf("❌ Error creating file: %v\n", err)
		return
	}
	// defer will close the file automatically when main ends
	defer file.Close()

	// Write pretty JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(users); err != nil {
		fmt.Printf("❌ Error writing JSON: %v\n", err)
		return
	}

	fmt.Println("✅ JSON written successfully to users.json")
	fmt.Println("   (Open the file to see the array of objects)")

	// =====================================================
	// 5. CONCURRENCY (GOROUTINES), CONTEXT, DEFER
	// =====================================================
	fmt.Println("\n=== 5. CONCURRENCY, CONTEXT & DEFER ===")

	var wg sync.WaitGroup
	// Create a context with 2-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel to release resources automatically
	defer cancel()

	// Example of defer (will run at the very end of main)
	defer fmt.Println("🧹 Main function cleanup via defer completed!")

	wg.Add(3) // 3 tasks

	// Goroutine 1 - simple task
	go func() {
		defer wg.Done()
		fmt.Println("🔄 Goroutine 1: Starting heavy calculation...")
		time.Sleep(500 * time.Millisecond) // simulate work
		fmt.Println("✅ Goroutine 1: Finished!")
	}()

	// Goroutine 2 - uses context for cancellation
	go func() {
		defer wg.Done()
		fmt.Println("🔄 Goroutine 2: Waiting for context or timeout...")
		select {
		case <-ctx.Done():
			fmt.Println("⏹️  Goroutine 2: Cancelled by context timeout!")
		case <-time.After(3 * time.Second):
			fmt.Println("✅ Goroutine 2: Completed normally")
		}
	}()

	// Goroutine 3 - demonstrates defer inside goroutine
	go func() {
		defer wg.Done()
		defer fmt.Println("🧹 Goroutine 3: Internal defer cleanup executed!")
		fmt.Println("🔄 Goroutine 3: Doing file-like work...")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("✅ Goroutine 3: Work done!")
	}()

	fmt.Println("🚀 Launched 3 goroutines concurrently...")
	wg.Wait() // wait for all goroutines
	fmt.Println("✅ All concurrent tasks finished!")

	// =====================================================
	// 6. GENERICS + POLYMORPHISM
	// =====================================================
	fmt.Println("\n=== 6. GENERICS & POLYMORPHISM ===")

	// Generics in action
	sumInts := Sum(10, 20, 30, 40)
	sumFloats := Sum(1.1, 2.2, 3.3, 4.4)
	fmt.Printf("Generic Sum[int]: %d\n", sumInts)
	fmt.Printf("Generic Sum[float64]: %.2f\n", sumFloats)

	// Polymorphism via interface (same code works for different types)
	speakers := []Speaker{
		Person{Name: "Alice"},
		Robot{Model: "X-3000"},
	}

	fmt.Println("Polymorphic calls (interfaces):")
	for _, s := range speakers {
		fmt.Println("   →", s.Speak())
	}

	// =====================================================
	// 7. CONTROL STRUCTURES (already shown in dedicated function)
	// =====================================================
	demonstrateControlStructures()

	fmt.Println("\n🎉 DEMO COMPLETE!")
	fmt.Println("✅ Check the created file: users.json")
	fmt.Println("✅ Run 'go run main.go' again anytime to see everything in action.")
	fmt.Println("GoLang features demonstrated: project init, types, collections, JSON I/O,")
	fmt.Println("concurrency + context + defer, generics, polymorphism, and all control structures.")
}
