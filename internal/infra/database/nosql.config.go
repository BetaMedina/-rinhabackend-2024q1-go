package infra

import (
	"context"
	"rinha/pkg"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(pkg.GetEnvs("MONGO_CONNECTION_STRING")))
	print(pkg.GetEnvs("MONGO_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	Client = client
}

func GetConnection(collection string) *mongo.Collection {
	return (*mongo.Collection)(Client.Database(pkg.GetEnvs("DATABASE")).Collection(collection))
}
