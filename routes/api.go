package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/yomraliahmet/fiber-api/controllers"
	_ "github.com/yomraliahmet/fiber-api/docs"
	"github.com/yomraliahmet/fiber-api/resources"
	"github.com/yomraliahmet/fiber-api/swagger"
)

func welcomeApi(c *fiber.Ctx) error {
	return c.Render("index", c)
}

func SetupApiRoutes(app *fiber.App) {

	api := app.Group("/api") // /api

	// Version 1
	v1 := api.Group("/v1.0") // /api/v1
	apiVersion1(v1)
}

func JwtError(c *fiber.Ctx, e error) error {
	response := resources.JsonResult(401, e.Error(), nil)
	return c.Status(401).JSON(response)
}

func apiVersion1(api fiber.Router) {

	api.Get("/docs/*", swagger.Handler)

	// welcome api endpoint
	api.Get("/", welcomeApi)

	api.Post("/auth/login", controllers.Login)

	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: JwtError,
	}))

	// User endpoints
	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:id", controllers.GetUser)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)

	// Product endpoints
	api.Post("/products", controllers.CreateProduct)
	api.Get("/products", controllers.GetProducts)
	api.Get("/products/:id", controllers.GetProduct)
	api.Put("/products/:id", controllers.UpdateProduct)
	api.Delete("/products/:id", controllers.DeleteProduct)

	// Order endpoints
	api.Post("/orders", controllers.CreateOrder)
	api.Get("/orders", controllers.GetOrders)
	api.Get("/orders/:id", controllers.GetOrder)
}
