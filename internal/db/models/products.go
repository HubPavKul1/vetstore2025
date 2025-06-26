package models


// Категория товара
type Category struct {
    BaseModel
	SubCategories []SubCategory
    
}

// Подкатегория товара
type SubCategory struct {
    BaseModel
    CategoryID 	uint
	Products []Product
}

// Товар номенклатура
type Product struct {
	BaseModel
	SubCategoryID uint
	ProductsInStore []ProductInStore
	PackagingID uint
	UnitID uint
}

// Наименование упаковки товара
type Packaging struct {
	BaseModel
	Products []Product
}

// Наименование единиц учета
type Unit struct {
	BaseModel
	Products []Product
}






