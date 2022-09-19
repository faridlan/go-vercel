package app

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection() (*mongo.Database, error) {

	pass := os.Getenv("db_pass")
	clientOptions := options.Client()
	clientOptions.ApplyURI(fmt.Sprintf("mongodb+srv://faridlan:%s@nostracode.oa4zwqi.mongodb.net", pass))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database("belajarMongo"), nil
}
