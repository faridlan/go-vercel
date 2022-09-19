package handler

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Handler(writer http.ResponseWriter, request *http.Request) {

	name := os.Getenv("name")

	fmt.Fprintf(writer, "Hello, %s", name)
}
