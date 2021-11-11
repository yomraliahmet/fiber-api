package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/template/html"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yomraliahmet/fiber-api/config"
	"github.com/yomraliahmet/fiber-api/database"
	"github.com/yomraliahmet/fiber-api/routes"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email fiber@fiber.com
// @host gofiberapi.herokuapp.com
// @BasePath /api/v1.0/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @accept json
// @produce json

func main() {

	config.LoadEnv()
	port := os.Getenv("PORT")

	database.ConnectDb()
	database.AutoMigrate()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Default config
	app.Use(cors.New())
	/*
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://127.0.0.1:8080, http://localhost:8080",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	*/
	routes.SetupWebRoutes(app)
	routes.SetupApiRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
