package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	test_db       = ""
	test_mongourl = ""
)

var client *mongo.Client

func Init() error {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(test_mongourl))
	if err != nil {
		return err
	}
	client = c
	return nil
}
