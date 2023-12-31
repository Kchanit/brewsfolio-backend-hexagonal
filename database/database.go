package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Kchanit/brewsfolio-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDb() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	fmt.Println("Connecting to database...")
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		os.Exit(2)
	}

	fmt.Println("DB connected")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Beer{})
	db.AutoMigrate(&models.Review{})
	db.AutoMigrate(&models.Collection{})
	DBConn = db
}

func IsUserEmpty() bool {
	var user []models.User
	DBConn.Find(&user)
	return len(user) == 0
}
