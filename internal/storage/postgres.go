package storage

import (
	"fmt"
	"log"

	"github.com/ZiplEix/API_template/internal/todo"
	"github.com/ZiplEix/API_template/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func BootstrapPostgres(owner, passwrd, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable", owner, passwrd, name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Print("failed to connect database\n", err)
		return nil, err
	}

	log.Println("Connected to database.")
	db.Logger = db.Logger.LogMode(logger.Info)

	log.Println("Migrating database...")
	err = db.AutoMigrate(todo.TodoDB{}, user.UserDB{})
	if err != nil {
		log.Fatal("Failed to migrate database.\n", err)
		return db, err
	}
	log.Println("Database migrated.")

	return db, nil
}

func ClosePostgres(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		return err
	}
	return database.Close()
}
