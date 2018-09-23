package models

// PRODUCT MODELS

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Global variable
var db *gorm.DB

// Product struct
type Product struct {
	gorm.Model
	ProductCode        string
	ProductName        string
	ProductDescription string
}

// InitialMigration does init migration when table in database doesn't exisr
func InitialMigration() {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})
}

// DatabaseConnection function does init connection with database
func DatabaseConnection() {
	db, err := gorm.Open("mysql", "root:reza@/store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
