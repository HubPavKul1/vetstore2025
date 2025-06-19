package models

import (
	"gorm.io/gorm"
)

// Категория товара
type Category struct {
    gorm.Model
    Name      string
	SubCategories []SubCategory
    
}

// Подкатегория товара
type SubCategory struct {
    gorm.Model
    Name        string
    CategoryID 	uint
	Products []Product
}

// Товар номенклатура
type Product struct {
	gorm.Model
	Name string
	SubCategoryID uint
	ProductsInStore []ProductInStore
	PackagingID uint
	UnitID uint
}

// Наименование упаковки товара
type Packaging struct {
	gorm.Model
	Name string
	Products []Product
}

// Наименование единиц учета
type Unit struct {
	gorm.Model
	Name string
	Products []Product
}






