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
