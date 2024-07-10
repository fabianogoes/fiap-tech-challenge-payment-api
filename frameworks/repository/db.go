package repository

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitDB(ctx context.Context, config *entities.Config) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.DBUri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	fmt.Println("Successfully connected to MongoDB")

	return client.Database(config.DBName), nil
}
