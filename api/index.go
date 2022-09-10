package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	w.Header().Add("content-type", "application/json")
	faridlan := User{
		Id:   1,
		Name: "Faridlan",
		Age:  21,
	}

	encode := json.NewEncoder(w)
	err := encode.Encode(&faridlan)
	if err != nil {
		panic(err)
	}
}
