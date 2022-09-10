package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
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

func Test(t *testing.T) {

	db := Connection()

	SQL := "SELECT id,name,age FROM users"
	rows, err := db.QueryContext(context.Background(), SQL)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Println(user)
	}
}
