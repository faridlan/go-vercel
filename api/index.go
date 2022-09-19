package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/faridlan/go-vercel/api/app"
	"github.com/faridlan/go-vercel/api/helper"
	"github.com/faridlan/go-vercel/api/model"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
)

func Handler(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Add("content-type", "application/json")
	user := FindAll(request.Context())

	encode := json.NewEncoder(writer)
	err := encode.Encode(&user)
	helper.FatalIfError(err)
}

func FindAll(ctx context.Context) []model.User {

	db, err := app.NewConnection()
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
