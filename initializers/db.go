package initializers

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/ilyasbabu/go-ecommerce/models"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("ecom.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting DB")
	}
	db.AutoMigrate(
		&models.Products{},
	)
}
