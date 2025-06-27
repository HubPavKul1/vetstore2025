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

func GetCategoryIDBYNameService(catName string) (uint, error) {
    return repository.CetCategoryIDByName(db.DB, catName)
}

// Подкатегории
func CreateSubCategoryService(subCat models.SubCategory) (*models.SubCategory, error) {
    return repository.CreateSubCategory(db.DB, subCat)
}

func GetSubCategoriesService() ([]models.SubCategory, error) {
    return repository.GetSubCategories(db.DB)
}


func GetSubCategoriesForCategoryService(categoryName string) ([]models.SubCategory, error) {
    return repository.GetSubCategoriesForCategory(db.DB, categoryName)
}

func GetSubCategoryIDBYNameService(subcatName string) (uint, error) {
    return repository.CetSubCategoryIDByName(db.DB, subcatName)
}
// Товары (номенклатура)
func CreateProductService(prod models.Product) (*models.Product, error) {
    return repository.CreateProduct(db.DB, prod)
}

func GetProductsService() ([]models.Product, error) {
    return repository.GetProducts(db.DB)
}

func GetProductsForSubCategoryService(subcategoryName string) ([]models.Product, error) {
    return repository.GetProductsForSubCategory(db.DB, subcategoryName)
}

//Упаковка
func CreatePackagingService(p models.Packaging) (*models.Packaging, error) {
    return repository.CreatePackaging(db.DB, p)
}

func GetPackagingService() ([]models.Packaging, error) {
    return repository.GetPackagings(db.DB)
}

func GetPackagingIDBYNameService(packName string) (uint, error) {
    return repository.CetPackagingIDByName(db.DB, packName)
}

//Единицы учета
func CreateUnitService(prod models.Unit) (*models.Unit, error) {
    return repository.CreateUnit(db.DB, prod)
}

func GetUnitsService() ([]models.Unit, error) {
    return repository.GetUnits(db.DB)
}

func GetUnitIDBYNameService(unitName string) (uint, error) {
    return repository.CetUnitIDByName(db.DB, unitName)
}