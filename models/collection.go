package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"user"`
	Beers       []Beer `gorm:"many2many:collection_beers" json:"beers"`
}
