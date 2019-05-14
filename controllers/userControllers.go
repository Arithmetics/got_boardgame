package controllers

import (
	"net/http"

	"github.com/arithmetics/got_boardgame/models"
	u "github.com/arithmetics/got_boardgame/utils"
)

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "*")
	(w).Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "OPTIONS" {
		return
	}

	id := r.Context().Value("user").(uint)
	data := models.GetUser(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
