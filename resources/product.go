package resources

import "github.com/yomraliahmet/fiber-api/models"

type Product struct {
	ID           uint   `json:"id" extensions:"x-order=1" example:"1"`
	Name         string `json:"name" extensions:"x-order=1" example:"Product One"`
	SerialNumber string `json:"serial_number" extensions:"x-order=1" example:"7551561457874455"`
}

func ProductResource(productModel models.Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}
