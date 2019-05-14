package models

import "github.com/jinzhu/gorm"

// Bid (5 power for the Fiefdoms track)
type Bid struct {
	gorm.Model
	TrackID   uint
	FactionID uint
	Amount    int
}
