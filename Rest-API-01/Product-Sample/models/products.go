package models

type Product struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

// Create Product
func (r *Product) CreateProduct(products *[]Product) (err error) {
	*products = append(*products, *r)
	return nil
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
func UpdateProductById(products *[]Product, NewProduct Product) (product Product) {
	for index, item := range *products {
		if item.Id == product.Id {
			(*products)[index].Name = product.Name
			(*products)[index].Price = product.Price
			(*products)[index].Stock = product.Stock

			return (*products)[index]
		}
	}
	return Product{}
}

// Delete The Product By Id
func (r *Product) DeleteProductById(products *[]Product) (product Product) {
	for index, item := range *products {
		if item.Id == (*r).Id {
			n := len(*products)
			deletedProduct := item
			(*products)[index] = (*products)[n-1]
			*products = (*products)[:n-1]
			return deletedProduct
		}
	}
	return Product{}
}
