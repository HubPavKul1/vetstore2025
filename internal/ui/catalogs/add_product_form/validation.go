package add_product_form

import (

	"fyne.io/fyne/v2/canvas"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

type AddProductFormField struct {
	FieldValue string
	FieldError *canvas.Text
}

func IsAddProductFormValid(
	category *AddProductFormField, 
	subcategory *AddProductFormField, 
	productName *AddProductFormField, 
	packName *AddProductFormField, 
	unitName *AddProductFormField,
) bool {
	valid := true
      
    if !ui_utils.IsValidSelect(category.FieldValue) {
        valid = false
        category.FieldError.Text = ui_utils.EmptyFieldError
        return valid
    }
        
    if !ui_utils.IsValidSelect(subcategory.FieldValue) {
        valid = false
        subcategory.FieldError.Text = ui_utils.EmptyFieldError
        return valid
    }
    
    if !ui_utils.IsNotEmptyField(productName.FieldValue) {
        valid = false
        productName.FieldError.Text = ui_utils.EmptyFieldError
        return valid
    }

    if !ui_utils.IsValidSelect(packName.FieldValue) {
        valid = false
        packName.FieldError.Text = ui_utils.EmptyFieldError
        return valid
    }
      
    if !ui_utils.IsValidSelect(unitName.FieldValue) {
        valid = false
        unitName.FieldError.Text = ui_utils.EmptyFieldError
        return valid
    }

    return valid
}

