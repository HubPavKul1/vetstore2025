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
	Phone string
	ProductReceipts []ProductReceipt
}

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
	PackagingAmount int
	UnitAmount float64
	Cost float64

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



