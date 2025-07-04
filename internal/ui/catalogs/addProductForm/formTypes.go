package addProductForm

import (
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
)



type AddProductFormData struct {
	SubcategoryID uint
	ProductName string
	PackagingID uint
	UnitID uint
}


type AddProductForm struct {
	CategorySelect *uiTypes.SelectWithError
	SubcategorySelect *uiTypes.SelectWithError
	ProductNameEntry *uiTypes.InputWithError
	PackagingSelect *uiTypes.SelectWithError
	UnitSelect *uiTypes.SelectWithError
	SaveButton *widget.Button
	BackButton *widget.Button
	
}