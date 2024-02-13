package repositories

import (
	"context"
	"rinha/internal/entities"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StatementRepository interface {
	Create(statement *entities.Statement) error
	List(id string) *[]entities.Statement
}
type statementRepository struct {
	collection *mongo.Collection
}

func (collection *statementRepository) Create(statement *entities.Statement) error {
	_, err := collection.collection.InsertOne(context.TODO(), bson.M{
		"client": bson.M{
			"_id":    statement.Client.ID,
			"limit":  statement.Client.Limite,
			"amount": statement.Client.Saldo,
			"id":     statement.Client.FriendlyId,
		},
		"date":        statement.Data,
		"description": statement.Descricao,
		"type":        statement.Tipo,
		"value":       statement.Valor,
	})
	if err != nil {
		return err
	}
	return err
}

func (collection *statementRepository) List(id string) *[]entities.Statement {
	var statements []entities.Statement
	formattedId, _ := strconv.Atoi(id)
	cursor, _ := collection.collection.Find(context.TODO(), bson.D{{Key: "client.id", Value: formattedId}}, options.Find().SetSort(bson.D{{Key: "date", Value: -1}}), options.Find().SetLimit(10))
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var statement entities.Statement
		if err := cursor.Decode(&statement); err != nil {
			return nil
		}
		statements = append(statements, statement)
	}
	return &statements

}

func NewStatementRepository(collection *mongo.Collection) StatementRepository {
	return &statementRepository{collection}
}
