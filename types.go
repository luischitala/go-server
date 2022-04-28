package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

//Define how the struct will behave with a specific parser
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//How to correctly parse to text the user model

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

//Type interface to handler de decoder

type MetaData interface{}
