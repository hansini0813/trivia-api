package database

import (
	"fmt"
	"log"
	"os"

	"github.com/hansini0813/trivia-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	// Build the Data Source Name (DSN) string for the Postgres connection
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		// Log the error and exit if the connection fails
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	// Log successful connection
	log.Println("Connected to the database")

	// Set the logger for the database connection
	db.Logger = logger.Default.LogMode(logger.Info)

	// Run the migrations for the `Fact` model
	log.Println("Running migrations")
	db.AutoMigrate(&models.Fact{})

	// Assign the database connection to the global `DB` instance
	DB = Dbinstance{
		Db: db,
	}
}
