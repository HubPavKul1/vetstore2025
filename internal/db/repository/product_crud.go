package repository

import (
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"gorm.io/gorm"
)

// Категории
func CreateCategory(db *gorm.DB, category models.Category) (*models.Category, error) {
    result := db.Create(&category)
    return &category, result.Error
}

func GetCategories(db *gorm.DB) ([]models.Category, error) {
    var categories []models.Category
    result := db.Find(&categories)
    return categories, result.Error
}

// Подкатегории
func CreateSubCategory(db *gorm.DB, subCat models.SubCategory) (*models.SubCategory, error) {
	result := db.Create(&subCat)
	return &subCat, result.Error
}

func GetSubCategories(db *gorm.DB) ([]models.SubCategory, error) {
	var subCategories []models.SubCategory
	result := db.Find(&subCategories)
	return subCategories, result.Error
}

// Упаковка
func CreatePackaging(db *gorm.DB, pack models.Packaging) (*models.Packaging, error) {
	result := db.Create(&pack)
	return &pack, result.Error
}

func GetPackagings(db *gorm.DB) ([]models.Packaging, error) {
	var packs []models.Packaging
	result := db.Find(&packs)
	return packs, result.Error
}

// Единицы учета
func CreateUnit(db *gorm.DB, unit models.Unit) (*models.Unit, error) {
	result := db.Create(&unit)
	return &unit, result.Error
}

func GetUnits(db *gorm.DB) ([]models.Unit, error) {
	var units []models.Unit
	result := db.Find(&units)
	return units, result.Error
}

// Товар
func CreateProduct(db *gorm.DB, prod models.Product) (*models.Product, error) {
	result := db.Create(&prod)
	return &prod, result.Error
}

func GetProducts(db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	result := db.Find(&products)
	return products, result.Error
}



