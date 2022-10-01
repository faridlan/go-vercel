package handler

import (
	"net/http"

	"github.com/faridlan/go-vercel/api/helper"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", helper.Hello)
}
