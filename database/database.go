package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Nivas-Mekala/ikea_inv_go_pg_docker/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectToDatabase() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pgPort := os.Getenv("DATABASE_PORT")
	pgHost := os.Getenv("DATABASE_HOST")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgName := os.Getenv("DB_NAME")
	pgTimeZone := os.Getenv("TIME_ZONE")

	//dsn := "dbname=postgres user=postgres password=password host=localhost port=5432 sslmode=disable TimeZone=Europe/Stockholm"
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s  port=%s sslmode=disable TimeZone=%s",
		pgName, pgUser, pgPassword, pgHost, pgPort, pgTimeZone)

	fmt.Println("Database details ::", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error Connecting to Database")
	}

	log.Println("Connect to data base")

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.Inventory{}, &models.Product{})

	Database = DBInstance{
		DB: db,
	}
}
