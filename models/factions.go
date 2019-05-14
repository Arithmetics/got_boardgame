package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Faction (Stark, Greyjoy, Lannister, etc.)
type Faction struct {
	gorm.Model
	Name        string
	PowerTokens int
	Game        Game
	GameID      uint
	User        User
	UserID      uint
	Bid         Bid
}

// MakeBid creates a new bid on the track
func (faction Faction) MakeBid(amount int, trackName string, db *gorm.DB) error {
	// bid cant be more than available power for faction
	if amount > faction.PowerTokens {
		return fmt.Errorf("Not enough power tokens")
	}

	var track Track
	db.Preload("Bids").Where(&Track{Name: trackName}).First(&track)

	if track.ID == 0 {
		return fmt.Errorf("no track found to bid on")
	}

	if !track.BiddingOpen {
		return fmt.Errorf("This track is not open for bidding")
	}
	// cant make two bids on the same track
	for _, bid := range track.Bids {
		if bid.FactionID == faction.ID {
			return fmt.Errorf("You already have a bid on this track")
		}
	}

	bid := Bid{
		TrackID:   track.ID,
		FactionID: faction.ID,
		Amount:    amount,
	}

	err := db.Create(&bid).Error

	return err
}
