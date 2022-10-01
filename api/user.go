package handler

import (
	"net/http"

	"github.com/faridlan/go-vercel/api/helper"
	"github.com/gorilla/mux"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {

	r := mux.NewRouter()

	r.HandleFunc("/api/user", helper.Hello).Methods("GET")
	r.HandleFunc("/", helper.Hello).Methods("GET")

	http.ListenAndServe(":80", r)
}
