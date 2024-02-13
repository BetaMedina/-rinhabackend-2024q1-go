package handlers

import (
	"rinha/internal/entities"
	"rinha/internal/repositories"
	"time"

	"github.com/gofiber/fiber"
)

type listStatement struct {
	clientRepository    repositories.ClientRepository
	statementRepository repositories.StatementRepository
}

type OutputDto struct {
	ID        string    `json:"id,omitempty"`
	Data      time.Time `json:"realizada_em" `
	Descricao string    `json:"descricao"`
	Tipo      string    `json:"tipo"`
	Valor     float64   `json:"valor"`
}

type ListStatement interface {
	executeTransaction(statement *[]entities.Statement) []OutputDto
	ListStatement(c *fiber.Ctx)
}

func (this *listStatement) executeTransaction(statements *[]entities.Statement) []OutputDto {
	var formattedStatement []OutputDto
	for _, statement := range *statements {
		formattedStatement = append(formattedStatement, OutputDto{
			Valor:     statement.Valor,
			Tipo:      statement.Tipo,
			Descricao: statement.Descricao,
			Data:      statement.Data,
		})
	}
	return formattedStatement
}

func (this *listStatement) ListStatement(c *fiber.Ctx) {
	result := this.statementRepository.List(c.Params("id"))
	countResults := len(*result)
	if countResults == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Client not found", "status": "error"})
		return
	}
	lastStatement := (*result)[countResults-1]
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"saldo": fiber.Map{
			"total":        lastStatement.Client.Saldo,
			"data_extrato": time.Now(),
			"limite":       lastStatement.Client.Limite,
		},
		"ultimas_transacoes": this.executeTransaction(result),
	})

}

func NewListStatement(clientRepository repositories.ClientRepository, statementRepository repositories.StatementRepository) ListStatement {
	return &listStatement{
		clientRepository:    clientRepository,
		statementRepository: statementRepository,
	}
}
