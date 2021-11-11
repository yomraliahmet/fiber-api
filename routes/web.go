package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func welcomeHome(c *fiber.Ctx) error {
	return c.Redirect("/api/v1.0")
}

func SetupWebRoutes(app *fiber.App) {
	// dashboard
	app.Get("/dashboard", monitor.New())

	// welcome home page endpoint
	app.Get("/", welcomeHome)
}
