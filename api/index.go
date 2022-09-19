package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/faridlan/go-vercel/api/helper"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
	Ages int                `bson:"ages,omitempty" json:"ages,omitempty"`
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Handler(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Add("content-type", "application/json")
	user := findAll(request.Context())

	encode := json.NewEncoder(writer)
	err := encode.Encode(&user)
	fatalIfError(err)
}

func findAll(ctx context.Context) []User {

	db, err := newConnection()
	fatalIfError(err)

	cursor, err := db.Collection("users").Find(ctx, bson.M{})
	fatalIfError(err)

	users := []User{}
	for cursor.Next(ctx) {
		var user User
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
