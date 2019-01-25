package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/arithmetics/auth_api/models"
	u "github.com/arithmetics/auth_api/utils"
)

//CreateGame makes a new game
var CreateGame = func(w http.ResponseWriter, r *http.Request) {

	creatorID := r.Context().Value("user").(uint) //Grab the id of the user that sent the request
	game := &models.Game{}

	err := json.NewDecoder(r.Body).Decode(game)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	game.Players = []uint{creatorID} //only player in game at begining is creator
	game.Creator = creatorID
	resp := game.Create()
	u.Respond(w, resp)
}

//GetMyGames finds the contracts for user
var GetMyGames = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetGamesCreatedBy(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
