package initializers

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/ilyasbabu/go-ecommerce/controllers"
	"github.com/ilyasbabu/go-ecommerce/models"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	var err error
	Db, err = gorm.Open(sqlite.Open("ecom.db"), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("Error Connecting DB")
	}
	err = Db.AutoMigrate(
		&models.Products{},
	)
	if err != nil {
		log.Print(err.Error())
		log.Fatal("failed to auto-migrate schema")
	}
	controllers.SetDB(Db)
}

func GetDB() *gorm.DB {
	return Db
}
