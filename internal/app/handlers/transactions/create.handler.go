package handlers

import (
	"errors"
	"rinha/internal/entities"
	"rinha/internal/repositories"
	"sync"
	"time"

	"github.com/gofiber/fiber"
)

type Body struct {
	Tipo      string  `json:"tipo"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}

type createTransaction struct {
	clientRepository    repositories.ClientRepository
	statementRepository repositories.StatementRepository
}

type CreateTransaction interface {
	exec(client entities.Client, row Body, wg *sync.WaitGroup)
	executeTransaction(client entities.Client, row Body) (entities.Client, error)
	CreateTransaction(c *fiber.Ctx)
}

func (this *createTransaction) exec(client entities.Client, row Body, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		this.clientRepository.Update(client.ID, client.Saldo)
	}()
	go func() {
		defer wg.Done()
		this.statementRepository.Create(&entities.Statement{
			Client: entities.Client{
				ID:         client.ID,
				FriendlyId: client.FriendlyId,
				Limite:     client.Limite,
				Saldo:      client.Saldo,
			},
			Descricao: row.Descricao,
			Data:      time.Now(),
			Tipo:      row.Tipo,
			Valor:     row.Valor,
		})
	}()
	wg.Wait()
}

func (this *createTransaction) executeTransaction(client entities.Client, row Body) (entities.Client, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	if row.Tipo == "c" {
		client.Saldo += row.Valor
		this.exec(client, row, &wg)
	}
	if row.Tipo == "d" {
		totalAmount := client.Saldo - row.Valor
		if totalAmount < -client.Limite {
			return entities.Client{}, errors.New("This amount is greater than your limit.")
		}
		client.Saldo = totalAmount
		row.Valor = -row.Valor
		this.exec(client, row, &wg)
	}
	return client, nil
}

func (this *createTransaction) CreateTransaction(c *fiber.Ctx) {
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "id is required",
		})
		return
	}
	client := this.clientRepository.FindClient(id)
	if client == nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Client not found",
		})
		return
	}
	body := new(Body)
	if err := c.BodyParser(body); err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "Payload ins't valid",
		})
		return
	}
	clientUpdated, err := this.executeTransaction(*client, *body)
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"limite": clientUpdated.Limite,
		"saldo":  clientUpdated.Saldo,
	})
	return

}
func NewCreateTransaction(clientRepository repositories.ClientRepository, statementRepository repositories.StatementRepository) CreateTransaction {
	return &createTransaction{clientRepository, statementRepository}
}
