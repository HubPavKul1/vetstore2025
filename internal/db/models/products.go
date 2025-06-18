package models

import (
	"time"

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

// Товар
type Product struct {
	gorm.Model
	Name string
	SubCategoryID uint
}

// Поставщики
type Supplier struct {
	gorm.Model
	Name string
	ProductReceipts []ProductReceipt

}

// Поступление товара
type ProductReceipt struct {
	gorm.Model
	ReceiptDate time.Time
	SupplierID uint
}

