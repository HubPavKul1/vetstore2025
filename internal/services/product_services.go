package service

import (
	"github.com/HubPavKul1/vetstore2025/internal/db"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/db/repository"
)

func CreateCategoryService(category models.Category) (*models.Category, error) {
    return repository.CreateCategory(db.DB, category)
}

func GetCategoriesService() ([]models.Category, error) {
    return repository.GetCategories(db.DB)
}

// Подкатегории
func CreateSubCategoryService(subCat models.SubCategory) (*models.SubCategory, error) {
    return repository.CreateSubCategory(db.DB, subCat)
}

func GetSubCategoriesService() ([]models.SubCategory, error) {
    return repository.GetSubCategories(db.DB)
}
