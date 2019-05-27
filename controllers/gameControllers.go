package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/arithmetics/got_boardgame/models"
	u "github.com/arithmetics/got_boardgame/utils"
)

// GetGame ...
func GetGame(w http.ResponseWriter, r *http.Request) {

}

// CreateGame ...
func CreateGame(w http.ResponseWriter, r *http.Request) {

	game := &models.Game{}
	err := json.NewDecoder(r.Body).Decode(game) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	userID := r.Context().Value("user").(uint)
	user := models.GetUserSimple(userID)

	game.UserCreator = userID
	game.Players = append(game.Players, *user)
	game.Active = true
	game.GameState = "waiting"
	resp := game.Create() //Create game
	u.Respond(w, resp)
}
