package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string       `json:"email"`
	Name        string       `json:"name"`
	Password    string       `json:"password"`
	Role        string       `json:"role"`
	Favorites   []Beer       `gorm:"many2many:user_beers" json:"favorites"`
	Reviews     []Review     `gorm:"foreignKey:UserID" json:"reviews"`
	Collections []Collection `gorm:"foreignKey:UserID" json:"collections"`
}
