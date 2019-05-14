package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arithmetics/got_boardgame/models"
	u "github.com/arithmetics/got_boardgame/utils"
)

// CreateContact creates a contact
func CreateContact(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserID = user
	resp := contact.Create()
	u.Respond(w, resp)
}

// GetContactsFor ...
func GetContactsFor(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// Test ii
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", "sdfdsf")
	resp := u.Message(true, "success")
	resp["data"] = "backend connected"
	u.Respond(w, resp)
}
