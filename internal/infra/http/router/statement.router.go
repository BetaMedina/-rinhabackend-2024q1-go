package router

import (
	handlers "rinha/internal/app/handlers/statement"
	infra "rinha/internal/infra/database"
	"rinha/internal/repositories"

	"github.com/gofiber/fiber"
)

var statementHandler = handlers.NewListStatement(repositories.NewClientRepository(
	infra.GetConnection("member")),
	repositories.NewStatementRepository(infra.GetConnection("statement")),
)

func StatementRouter(app *fiber.App) {
	app.Get("/clientes/:id/extrato", statementHandler.ListStatement)
}
