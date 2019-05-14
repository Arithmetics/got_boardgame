package models

import (
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
)

//Track is the power track (Sword, Throne, Raven)
type Track struct {
	gorm.Model
	Name        string
	Game        Game
	GameID      uint
	Bids        []Bid
	BiddingOpen bool
	Position1   Faction
	Position1ID uint
	Position2   Faction
	Position2ID uint
	Position3   Faction
	Position3ID uint
	Position4   Faction
	Position4ID uint
	Position5   Faction
	Position5ID uint
	Position6   Faction
	Position6ID uint
}

// BiddingComplete is true if all players have placed a bid
func (track Track) BiddingComplete() bool {
	if len(track.Bids) < 6 {
		return false
	}

	return true
}

// SettleBids ends bidding, fills the positions, and adjusts everyones power tokens
func (track Track) SettleBids(db *gorm.DB) error {
	db.Preload("Bids").Find(&track)

	if !track.BiddingComplete() {
		return fmt.Errorf("Track is not ready for bidding to end")
	}

	sort.Sort(byAmount(track.Bids))

	track.Position1ID = track.Bids[0].FactionID
	track.Position2ID = track.Bids[1].FactionID
	track.Position3ID = track.Bids[2].FactionID
	track.Position4ID = track.Bids[3].FactionID
	track.Position5ID = track.Bids[4].FactionID
	track.Position6ID = track.Bids[5].FactionID

	err := db.Save(&track).Error

	if err != nil {
		return err
	}

	for _, bid := range track.Bids {
		var faction Faction
		db.Find(&faction, bid.FactionID)

		faction.PowerTokens = faction.PowerTokens - bid.Amount

		db.Save(&faction)

		db.Delete(&bid)
	}

	return nil

}

// OpenBidding starts bidding track
func (track Track) OpenBidding(db *gorm.DB) {
	track.BiddingOpen = true

	db.Save(&track)
}

// CloseBidding closes bidding on the track
func (track Track) CloseBidding(db *gorm.DB) {
	track.BiddingOpen = true

	db.Save(&track)
}

type byAmount []Bid

func (s byAmount) Len() int {
	return len(s)
}
func (s byAmount) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byAmount) Less(i, j int) bool {
	return s[i].Amount > s[j].Amount
}
