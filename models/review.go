package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	BeerID      uint   `gorm:"not null" json:"beer_id"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"user"`
	Rating      uint   `gorm:"not null" json:"rating"`
	Description string `json:"description"`
}
