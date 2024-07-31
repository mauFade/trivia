package db

import (
	"fmt"
	"log"
	"os"

	"github.com/mauFade/trivia/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

var Database DatabaseInstance

func OpenDatabase() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error connecting to database.\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database!")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")

	db.AutoMigrate(
		&models.Fact{},
	)

	Database = DatabaseInstance{
		DB: db,
	}
}
