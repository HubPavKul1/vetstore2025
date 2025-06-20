package repository

import (
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) (*models.Category, error) {
    result := db.Create(&category)
    return &category, result.Error
}

func GetCategories(db *gorm.DB) ([]models.Category, error) {
    var categories []models.Category
    result := db.Find(&categories)
    return categories, result.Error
}



