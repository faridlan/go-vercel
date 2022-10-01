package handler

import (
	"fmt"
	"net/http"

	"github.com/faridlan/go-vercel/api/helper"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	result := helper.Hello()
	fmt.Fprintf(writer, result)
}
