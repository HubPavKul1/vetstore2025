package models

import (
	"math"
	"time"

	"gorm.io/gorm"
)

// Поступление товара
type ProductReceipt struct {
	gorm.Model
	ReceiptDate time.Time
	SupplierID uint
	ProductsInStore []ProductInStore
}

// Поступивший товар
type ProductInStore struct {
	gorm.Model
	ProductID uint
	ProductReceiptID uint
	PackagingID uint
	PackagingAmount int
	UnitID uint
	UnitAmount float64
	Cost float64

}
// Метод расчета цены товара
func(prod ProductInStore) GetPrice() float64 {
	price := prod.Cost / prod.UnitAmount

	return math.Round(price * 100) / 100
}

// Перемещение товара в подразделения
type MovingFromStore struct {
	gorm.Model
	MovingDate time.Time
	DepartmentID uint
}



