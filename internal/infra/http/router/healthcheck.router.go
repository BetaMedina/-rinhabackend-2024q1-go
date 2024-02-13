package router

import "github.com/gofiber/fiber"

func HealthCheckRouter(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"status": "success",
			"data":   "Let's go! 🚀",
		})
	})
}
