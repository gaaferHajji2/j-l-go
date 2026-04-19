import (
	"encoding/json"
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