package db

import (
	"fmt"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"

	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

// InitDB открывает соединение с базой данных и возвращает ошибку, если возникнут проблемы.
func InitDB(filePath string) error {
    var err error
    DB, err = gorm.Open(sqlite.Open(filePath), &gorm.Config{})
    if err != nil {
        return fmt.Errorf("failed to connect database: %w", err)
    }

    // Выполняем миграцию
    err = migrate(DB)
    if err != nil {
        return fmt.Errorf("migration failed: %w", err)
    }

    fmt.Println("Database connected successfully.")
    return nil
}

// Migrate выполняет миграцию базы данных.
func migrate(db *gorm.DB) error {
    err := db.AutoMigrate(
        &models.Category{},
        &models.SubCategory{},
        &models.Product{},
        &models.Packaging{},
        &models.Unit{},
        //Склад
        &models.Supplier{},
        &models.ProductReceipt{},
        &models.ProductInStore{},
        &models.MovingFromStore{},
        // Подразделения
        &models.Department{},
        &models.Position{},
        &models.Employee{},
        
       

     )
    if err != nil {
        return fmt.Errorf("auto migration failed: %w", err)
    }
    return nil
}

// Close закрывает установленное ранее соединение с базой данных.
func Close() error {
    if DB != nil {
        sqlDB, err := DB.DB() // Получаем raw sql.DB
        if err != nil {
            return err
        }
        return sqlDB.Close() // Закрываем raw sql.DB
    }
    return nil
}