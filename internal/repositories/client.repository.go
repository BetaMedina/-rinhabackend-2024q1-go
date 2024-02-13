package repositories

import (
	"context"
	"rinha/internal/entities"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository interface {
	FindClient(id string) *entities.Client
	Update(id string, amount float64) error
}
type clientRepository struct {
	collection *mongo.Collection
}

func (collection *clientRepository) FindClient(id string) *entities.Client {
	var client entities.Client
	clientID, _ := strconv.Atoi(id)
	collection.collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: clientID}}).Decode(&client)
	return &client
}
func (collection *clientRepository) Update(id string, amount float64) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.collection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}, bson.D{{Key: "$set", Value: bson.D{
		{Key: "amount", Value: amount},
	}}})

	if err != nil {
		return err
	}
	return nil
}
func NewClientRepository(collection *mongo.Collection) ClientRepository {
	return &clientRepository{collection}
}
