package models

import (
	"time"

	"github.com/HubPavKul1/vetstore2025/internal/utils"
	"gorm.io/gorm"
)

// Поставщики
type Supplier struct {
	gorm.Model
	Name string
	PhoneNumber string `gorm:"null"`
	ProductReceipts []ProductReceipt
}

// Поступление товара
type ProductReceipt struct {
	gorm.Model
	ReceiptDate time.Time
	SupplierID uint
	InvoiceNumber string `gorm:"null"`
	InvoiceFile []byte `gorm:"null"`
	ProductsInStore []ProductInStore
}

// Поступивший товар
type ProductInStore struct {
	gorm.Model
	ProductID uint
	ProductReceiptID uint
	PackagingAmount int `gorm:"default:0"`
	UnitAmount float64	`gorm:"default:0.0"`
	Cost float64 `gorm:"default:0.0"`

}
// Метод расчета цены товара
func(prod ProductInStore) GetPrice() float64 {
	price := prod.Cost / prod.UnitAmount

	return utils.RoundFloat(price, 2)
}

// Перемещение товара в подразделения
type MovingFromStore struct {
	gorm.Model
	MovingDate time.Time
	DepartmentID uint
	
}



