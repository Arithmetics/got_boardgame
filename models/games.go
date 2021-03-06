package models

import (
	"fmt"
	"math/rand"
	"time"

	u "github.com/arithmetics/got_boardgame/utils"
	"github.com/jinzhu/gorm"
)

//Game struct (FirstGame, SecondGame, etc...)
type Game struct {
	gorm.Model
	UserCreator uint      `json:"userCreator"`
	Players     []User    `gorm:"many2many:joined_games" json:"players"`
	Name        string    `json:"name"`
	Active      bool      `json:"active"`
	Factions    []Faction `json:"factions"`
	Tracks      []Track   `json:"tracks"`
	GameState   string    `json:"gameState"`
}

//Validate incoming game details..
func (game *Game) Validate() (map[string]interface{}, bool) {

	if len(game.Name) < 1 {
		return u.Message(false, "Game name not long enough"), false
	}

	return u.Message(false, "Requirement passed"), true
}

// Create makes a new game in the db
func (game *Game) Create() map[string]interface{} {

	if resp, ok := game.Validate(); !ok {
		return resp
	}

	GetDB().Create(game)

	if game.ID <= 0 {
		return u.Message(false, "Failed to create game, connection error.")
	}

	if err := game.CreateTracks(); err != nil {
		return u.Message(false, "Error creating tracks")
	}

	response := u.Message(true, "Game has been created")
	response["game"] = game
	return response
}

// ClearPasswords clears passwords that may be included for a game players
func (game *Game) ClearPasswords() {
	for i := range game.Players {
		game.Players[i].ClearPassword()
	}
}

// GetGame grabs a game by ID
func GetGame(u uint) *Game {
	game := &Game{}
	GetDB().Preload("Tracks").Preload("Factions").Preload("Players").Table("games").Where("id = ?", u).First(game)
	if game.Name == "" { //Game not found!
		return nil
	}

	for i := range game.Players {
		game.Players[i].Password = ""
		game.Players[i].Token = ""
	}

	return game
}

// AssignFactions creates a faction for each user in the game
func (game Game) AssignFactions(db *gorm.DB) error {

	if len(game.Players) < 6 {
		return fmt.Errorf("Not enough players to start this game")
	}

	factionNames := []string{"Stark", "Greyjoy", "Lannister", "Barratheon", "Martell", "Tyrell", "Bolton"}

	factionNames = shuffle(factionNames)

	for i, name := range factionNames {
		faction := Faction{
			Name:        name,
			PowerTokens: 10,
			GameID:      game.ID,
			User:        game.Players[i],
		}

		db.Create(&faction)
	}

	return nil
}

// CreateTracks is part of the start up sequernce. Creates tracks for a game in their default state
func (game Game) CreateTracks() error {
	track1 := Track{
		Name:        "IronThrone",
		GameID:      game.ID,
		BiddingOpen: true,
	}
	track2 := Track{
		Name:        "Fiefdoms",
		GameID:      game.ID,
		BiddingOpen: false,
	}
	track3 := Track{
		Name:        "KingsCourt",
		GameID:      game.ID,
		BiddingOpen: false,
	}

	GetDB().Create(&track1)
	GetDB().Create(&track2)
	GetDB().Create(&track3)

	return nil
}

//OpenTrackBidding allows bidding to begin on a track
func (game Game) OpenTrackBidding(trackname string, db *gorm.DB) error {
	return nil
}

func shuffle(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

// DeleteGame deletes a game by ID
func DeleteGame(gameID uint, deleterID uint) bool {
	game := &Game{}
	db := GetDB()
	db.Table("games").Where("id = ?", gameID).First(game)
	if game.Name == "" { //Game not found!
		return false
	}

	if game.UserCreator != deleterID {
		return false
	}

	db.Delete(game)

	return true
}
