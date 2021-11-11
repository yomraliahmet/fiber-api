package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yomraliahmet/fiber-api/database"
	"github.com/yomraliahmet/fiber-api/models"
	"github.com/yomraliahmet/fiber-api/resources"
)

// Create User
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	database.Database.Db.Create(&user)

	response := resources.JsonResult(200, "New user created successfully.", resources.UserResource(user))

	return c.Status(200).JSON(response)

}

// Get Users
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)

	responseUsers := []resources.User{}

	for _, user := range users {
		responseUser := resources.UserResource(user)
		responseUsers = append(responseUsers, responseUser)
	}

	response := resources.JsonResult(200, "Successfully.", responseUsers)

	return c.Status(200).JSON(response)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

// Get user
func GetUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		response := resources.JsonResult(400, "Please ensure that :id is an integer.", nil)

		return c.Status(400).JSON(response)
	}

	if err := findUser(id, &user); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	response := resources.JsonResult(200, "Successfully.", resources.UserResource(user))

	return c.Status(200).JSON(response)
}

// Update User
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		response := resources.JsonResult(400, "Please ensure that :id is an integer.", nil)

		return c.Status(400).JSON(response)
	}

	if err := findUser(id, &user); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		City      string `json:"city"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		response := resources.JsonResult(500, err.Error(), nil)

		return c.Status(500).JSON(response)
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.City = updateData.City

	database.Database.Db.Save(&user)

	response := resources.JsonResult(200, "Successfully.", resources.UserResource(user))

	return c.Status(200).JSON(response)
}

// Delete User
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		response := resources.JsonResult(400, "Please ensure that :id is an integer.", nil)

		return c.Status(400).JSON(response)
	}

	if err := findUser(id, &user); err != nil {
		response := resources.JsonResult(400, err.Error(), nil)

		return c.Status(400).JSON(response)
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		response := resources.JsonResult(404, err.Error(), nil)

		return c.Status(404).JSON(response)
	}

	response := resources.JsonResult(200, "Successfully.", nil)

	return c.Status(200).JSON(response)
}

func Topla(a int, b int) int {
	return a + b
}
