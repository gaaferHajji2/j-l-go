package models

type Product struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}
