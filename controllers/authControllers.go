package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arithmetics/got_boardgame/models"
	u "github.com/arithmetics/got_boardgame/utils"
)

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Create user
	u.Respond(w, resp)
}

// Authenticate ...
func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", "yyyyyy")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(w.Header())
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}
