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
