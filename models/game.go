package models

import (
	"fmt"

	u "github.com/arithmetics/auth_api/utils"
	"github.com/jinzhu/gorm"
)

//Game is good
type Game struct {
	gorm.Model
	Creator  uint   `json:"creator_id"` //The user who started the game
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
	Players  []uint `json:"players"` //The user ids in the game
}

//Validate returns message and true if the requirement is met
func (game *Game) Validate() (map[string]interface{}, bool) {

	if game.Nickname == "" {
		return u.Message(false, "Game needs a nickname"), false
	}

	if game.Status == "" {
		return u.Message(false, "game needs a status"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

//Create  creates a contact on contact struct
func (game *Game) Create() map[string]interface{} {

	if resp, ok := game.Validate(); !ok {
		return resp
	}

	GetDB().Create(game)

	resp := u.Message(true, "success")
	resp["game"] = game
	return resp
}

//GetGame  finds game by ID
func GetGame(id uint) *Game {

	game := &Game{}
	err := GetDB().Table("games").Where("id = ?", id).First(game).Error
	if err != nil {
		return nil
	}
	return game
}

//GetGamesCreatedBy for a user ID
func GetGamesCreatedBy(creatorID uint) []*Game {

	games := make([]*Game, 0)
	err := GetDB().Table("games").Where("creator_id = ?", creatorID).Find(&games).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return games
}
