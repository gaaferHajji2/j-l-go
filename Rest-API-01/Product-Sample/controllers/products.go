package controllers

import "jloka/jloka/models"

type ProductRepo struct {
	Products *[]models.Product
}

func Init(products *[]models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}
