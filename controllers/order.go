package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yomraliahmet/fiber-api/database"
	"github.com/yomraliahmet/fiber-api/models"
	"github.com/yomraliahmet/fiber-api/resources"
)

// Create Order
func CreateOrder(c *fiber.Ctx) error {

	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	database.Database.Db.Create(&order)

	responseUser := resources.UserResource(user)
	responseProduct := resources.ProductResource(product)
	responseOrder := resources.OrderResource(order, responseUser, responseProduct)

	response := resources.JsonResult(200, "New order created successfully.", responseOrder)

	return c.Status(200).JSON(response)

}

// Get Orders
func GetOrders(c *fiber.Ctx) error {

	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []resources.Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		reponseOrder := resources.OrderResource(order, resources.UserResource(user), resources.ProductResource(product))
		responseOrders = append(responseOrders, reponseOrder)
	}

	response := resources.JsonResult(200, "Successfully.", responseOrders)

	return c.Status(200).JSON(response)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}

// Get Order
func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil {
		response := resources.JsonResult(400, "Please ensure that :id is an integer.", nil)

		return c.Status(400).JSON(response)
	}

	if err := findOrder(id, &order); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)

	responseUser := resources.UserResource(user)
	responseProduct := resources.ProductResource(product)

	responseOrder := resources.OrderResource(order, responseUser, responseProduct)

	response := resources.JsonResult(200, "Successfully.", responseOrder)

	return c.Status(200).JSON(response)
}
