package services

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

func GetSubCategoriesForCategory(categoryId uint) ([]models.SubCategory, error) {
    return repository.GetSubCategoriesForCategory(db.DB, categoryId)
}

// Товары (номенклатура)
func CreateProductService(prod models.Product) (*models.Product, error) {
    return repository.CreateProduct(db.DB, prod)
}

func GetProductsService() ([]models.Product, error) {
    return repository.GetProducts(db.DB)
}

func GetProductsForSubCategoryService(subcategoryId uint) ([]models.Product, error) {
    return repository.GetProductsForSubCategory(db.DB, subcategoryId)
}

//Упаковка
func CreatePackagingService(p models.Packaging) (*models.Packaging, error) {
    return repository.CreatePackaging(db.DB, p)
}

func GetPackagingService() ([]models.Packaging, error) {
    return repository.GetPackagings(db.DB)
}

//Единицы учета
func CreateUnitService(prod models.Unit) (*models.Unit, error) {
    return repository.CreateUnit(db.DB, prod)
}

func GetUnitsService() ([]models.Unit, error) {
    return repository.GetUnits(db.DB)
}