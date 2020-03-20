package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client      *mongo.Client
	user     *mongo.Collection
}

func New(uri string) (*Repository, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	user := client.Database("User").Collection("user")



	return &Repository{
		client:      client,
		user:user,
	}, nil
}
