package initializers

import (
	"log"
	"os"

	"github.com/ilyasbabu/go-ecommerce/controllers"
	"github.com/ilyasbabu/go-ecommerce/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DB_STRING")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("Error Connecting DB")
	}
}

func SyncDatabase() {
	err := Db.AutoMigrate(
		&models.Products{},
		&models.User{},
		&models.ProductImages{},
	)
	if err != nil {
		log.Print(err.Error())
		log.Fatal("failed to auto-migrate schema")
	}
	controllers.SetDB(Db)
	services.SetDB(Db)
}
