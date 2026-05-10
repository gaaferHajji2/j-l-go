package models

type Product struct {
	Id    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Stock int     `json:"stock" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

// Create Product
func (r *Product) CreateProduct(products *[]Product, product *Product) (result *[]Product, err error) {
	*products = append(*products, *product)

	return products, nil
}

// Get the Product By Id
func (r *Product) GetProductById(products *[]Product) (product Product) {
	for _, item := range *products {
		if item.Id == r.Id {
			return item
		}
	}
	return Product{}
}

// Update the Product By Id
func (r *Product) UpdateProductById(products *[]Product) (product Product) {
	for index, item := range *products {
		if item.Id == (*r).Id {
			(*products)[index].Name = (*r).Name
			(*products)[index].Price = (*r).Price
			(*products)[index].Stock = (*r).Stock

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
