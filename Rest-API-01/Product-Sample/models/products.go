package models

type Product struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

// Create Product
func CreateProduct(products *[]Product, product *Product) (result *[]Product, err error) {
	*products = append(*products, *product)

	return products, nil
}

// Get the Product By Id
func GetProductById(products *[]Product, id int) (product Product) {
	for _, item := range *products {
		if item.Id == id {
			return item
		}
	}
	return Product{}
}

// Update the Product By Id
func UpdateProductById(products *[]Product, newProduct *Product) (product Product) {
	for index, item := range *products {
		if item.Id == (*newProduct).Id {
			(*products)[index].Name = (*newProduct).Name
			(*products)[index].Price = (*newProduct).Price
			(*products)[index].Stock = (*newProduct).Stock

			return (*products)[index]
		}
	}
	return Product{}
}
