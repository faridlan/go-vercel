package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func Connection() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(128.199.67.28:3306)/go_vercel")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func getUser() ([]User, error) {
	db := Connection()

	SQL := "SELECT id,name,age FROM users"
	rows, err := db.QueryContext(context.Background(), SQL)
	if err != nil {
		panic(err)
	}

	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return users, err
		}

		users = append(users, user)

	}
	return users, nil
}

// func Handler(w http.ResponseWriter, r *http.Request) {

// 	user, err := getUser()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// faridlan := User{
// 	// 	Id:   1,
// 	// 	Name: "Faridlan",
// 	// 	Age:  21,
// 	// }

// 	w.Header().Add("content-type", "application/json")
// 	encode := json.NewEncoder(w)
// 	err = encode.Encode(&user)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func Hello(writer http.ResponseWriter, request *http.Request) {

	user, err := getUser()
	if err != nil {
		panic(err)
	}
	// faridlan := User{
	// 	Id:   1,
	// 	Name: "Faridlan",
	// 	Age:  21,
	// }

	writer.Header().Add("content-type", "application/json")
	encode := json.NewEncoder(writer)
	err = encode.Encode(&user)
	if err != nil {
		panic(err)
	}

}
