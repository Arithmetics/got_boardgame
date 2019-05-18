package models

import (
	"fmt"
	"math/rand"
	"time"

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

// AssignFactions creates a faction for each user in the game
func (game Game) AssignFactions(db *gorm.DB) error {

	if len(game.Players) < 6 {
		return fmt.Errorf("Not enough players to start this game")
	}

	factionNames := []string{"Stark", "Greyjoy", "Lannister", "Barratheon", "Martell", "Tyrell"}

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
func (game Game) CreateTracks(db *gorm.DB) error {
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

	db.Create(&track1)
	db.Create(&track2)
	db.Create(&track3)

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
