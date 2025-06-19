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
}

// Наименование упаковки товара
type Packaging struct {
	gorm.Model
	Name string
	ProductsInStore []ProductInStore
}

// Наименование единиц учета
type Unit struct {
	gorm.Model
	Name string
	ProductsInStore []ProductInStore
}

// Поставщики
type Supplier struct {
	gorm.Model
	Name string
	Phone string
	ProductReceipts []ProductReceipt
}




