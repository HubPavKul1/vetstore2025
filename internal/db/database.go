package db

import (
    "fmt"
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"models"
)

var DB *gorm.DB

func InitDB() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("vetstore2025.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect database: %v", err)
        return err
    }

    // Создаем таблицы
    err = migrate()
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
        return err
    }
    fmt.Println("Database connected successfully.")
    return nil
}

func migrate() error {
    err := DB.AutoMigrate(
		&models.Category{}, 
		&models.SubCategory{},
		) // Добавьте остальные модели
    if err != nil {
        return fmt.Errorf("migration failed: %w", err)
    }
    return nil
}