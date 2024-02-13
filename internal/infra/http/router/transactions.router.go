package router

import (
	handlers "rinha/internal/app/handlers/transactions"
	infra "rinha/internal/infra/database"
	"rinha/internal/repositories"

	"github.com/gofiber/fiber"
)

var transactionHandler = handlers.NewCreateTransaction(repositories.NewClientRepository(
	infra.GetConnection("member")),
	repositories.NewStatementRepository(infra.GetConnection("statement")),
)

func TransactionsRouter(app *fiber.App) {
	app.Post("/clientes/:id/transacoes", transactionHandler.CreateTransaction)
}
