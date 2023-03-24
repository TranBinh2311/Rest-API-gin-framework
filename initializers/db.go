package initializers

import (
	"log"
	"os"

	"github.com/example/gin_framework/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect Database, details error", err)
	}
}

func MigrateEnvVariables() {
	DB.AutoMigrate(&model.User{})
}
