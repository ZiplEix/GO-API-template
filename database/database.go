package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ZiplEix/API_template/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB *DBInstance

func ConnectDB() error {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable ",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		return err
	}

	log.Println("Connected to database.")
	db.Logger = db.Logger.LogMode(logger.Info)

	log.Println("Migrating database...")
	err = db.AutoMigrate(models.Todo{}, models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database.\n", err)
		return err
	}
	log.Println("Database migrated.")

	DB = &DBInstance{Db: db}

	return nil
}

func CloseDB() {
	db, err := DB.Db.DB()
	if err != nil {
		log.Fatal("Failed to close database.\n", err)
	}
	db.Close()
	log.Println("Database closed.")
}
