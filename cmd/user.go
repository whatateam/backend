package cmd

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Email     string
}

var users = []User{
	User{Id: 1, FirstName: "Michael", LastName: "Zion", UserName: "michaelzion", Password: "123", Email: "noisleahcim@gmail.com"},
}

var UsersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})
