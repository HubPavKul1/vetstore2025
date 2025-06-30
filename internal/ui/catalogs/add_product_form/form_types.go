package add_product_form

import (
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_types"
)



type AddProductFormData struct {
	SubcategoryID uint
	ProductName string
	PackagingID uint
	UnitID uint
}


type AddProductForm struct {
	CategorySelect *ui_types.SelectWithError
	SubcategorySelect *ui_types.SelectWithError
	ProductNameEntry *ui_types.InputWithError
	PackagingSelect *ui_types.SelectWithError
	UnitSElect *ui_types.SelectWithError
	SaveButton *widget.Button
	BackButton *widget.Button
	
}