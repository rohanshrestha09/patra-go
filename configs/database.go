package configs

import (
	"log"
	"os"

	"github.com/rohanshrestha09/patra-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_LOCAL")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	if err = db.AutoMigrate(&models.User{}, &models.Chat{}, &models.Message{}); err != nil {
		log.Fatal("Error while migrating")
	}

	return db

}
