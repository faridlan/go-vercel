package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/faridlan/go-vercel/api/helper"
	"github.com/faridlan/go-vercel/api/model"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Handler(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Add("content-type", "application/json")
	user := FindAll(request.Context())

	encode := json.NewEncoder(writer)
	err := encode.Encode(&user)
	helper.FatalIfError(err)
}

func FindAll(ctx context.Context) []model.User {

	db, err := newConnection()
	helper.FatalIfError(err)

	cursor, err := db.Collection("users").Find(ctx, bson.M{})
	helper.FatalIfError(err)

	users := []model.User{}
	for cursor.Next(ctx) {
		var user model.User
		err := cursor.Decode(&user)
		helper.FatalIfError(err)

		users = append(users, user)
	}

	return users
}

func newConnection() (*mongo.Database, error) {

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
