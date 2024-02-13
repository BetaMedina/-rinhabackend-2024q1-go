package main

import (
	"rinha/internal/infra/http/router"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	router.HealthCheckRouter(app)
	router.TransactionsRouter(app)
	router.StatementRouter(app)
	app.Listen(":8000")
}
