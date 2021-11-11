package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yomraliahmet/fiber-api/request"
	"github.com/yomraliahmet/fiber-api/resources"
)

func Login(c *fiber.Ctx) error {

	payload := request.Auth{}
	if err := c.BodyParser(&payload); err != nil {

		response := resources.JsonResult(fiber.StatusUnauthorized, "Unauthorized", nil)

		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Throws Unauthorized error
	if payload.User != "john" || payload.Pass != "doe" {

		response := resources.JsonResult(fiber.StatusUnauthorized, "Unauthorized", nil)

		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		response := resources.JsonResult(fiber.StatusInternalServerError, err.Error(), nil)

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := resources.JsonResult(200, "Successfully", fiber.Map{"token": t})

	return c.Status(200).JSON(response)

}
