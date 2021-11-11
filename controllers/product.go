package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yomraliahmet/fiber-api/database"
	"github.com/yomraliahmet/fiber-api/models"
	"github.com/yomraliahmet/fiber-api/resources"
)

// Create Product
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	database.Database.Db.Create(&product)

	response := resources.JsonResult(200, "New user created successfully.", resources.ProductResource(product))

	return c.Status(200).JSON(response)

}

// Get Products
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []resources.Product{}

	for _, product := range products {
		responseProduct := resources.ProductResource(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	response := resources.JsonResult(200, "Successfully.", responseProducts)

	return c.Status(200).JSON(response)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}

	return nil
}

// Get Product
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		response := resources.JsonResult(400, "Please ensure that :id is an integer.", nil)

		return c.Status(400).JSON(response)
	}

	if err := findProduct(id, &product); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	response := resources.JsonResult(200, "Successfully.", resources.ProductResource(product))

	return c.Status(200).JSON(response)
}

// Update Product
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		response := resources.JsonResult(400, "Please ensure that :id is an integer.", nil)

		return c.Status(400).JSON(response)
	}

	if err := findProduct(id, &product); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		response := resources.JsonResult(500, err.Error(), nil)

		return c.Status(500).JSON(response)
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	response := resources.JsonResult(200, "Successfully.", resources.ProductResource(product))

	return c.Status(200).JSON(response)
}

// Delete Product
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		response := resources.JsonResult(400, "Please ensure that :id is an integer.", nil)

		return c.Status(400).JSON(response)
	}

	if err := findProduct(id, &product); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		response := resources.JsonResult(404, err.Error(), nil)

		return c.Status(404).JSON(response)
	}

	response := resources.JsonResult(200, "Successfully.", nil)

	return c.Status(200).JSON(response)
}
