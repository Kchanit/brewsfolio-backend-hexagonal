package models

import "gorm.io/gorm"

type Beer struct {
	gorm.Model
	Name             string   `gorm:"not null" json:"name"`
	Tagline          string   `gorm:"not null" json:"tagline"`
	Description      string   `gorm:"not null" json:"description"`
	ImageURL         string   `gorm:"not null" json:"image_url"`
	Abv              float32  `gorm:"not null" json:"abv"`
	Ibu              float32  `json:"ibu"`
	Ebc              int      `json:"ebc"`
	Srm              float32  `json:"srm"`
	Ph               float32  `json:"ph"`
	AttenuationLevel float32  `gorm:"not null" json:"attenuation_level"`
	Reviews          []Review `gorm:"foreignKey:BeerID"`
	UserID           uint     `json:"user_id"`
}
