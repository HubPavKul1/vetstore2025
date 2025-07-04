package addProductForm

import (

	// "fyne.io/fyne/v2/canvas"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)


func IsAddProductFormValid(form *AddProductForm) bool {
	valid := true
      
    if !uiUtils.IsValidSelect(form.CategorySelect.Select.Selected) {
        valid = false
        form.CategorySelect.ErrorLabel.Text = uiUtils.EmptyFieldError
        return valid
    }
        
    if !uiUtils.IsValidSelect(form.CategorySelect.Select.Selected) {
        valid = false
        form.SubcategorySelect.ErrorLabel.Text = uiUtils.EmptyFieldError
        return valid
    }
    
    if !uiUtils.IsNotEmptyField(form.ProductNameEntry.Input.Text) {
        valid = false
        form.ProductNameEntry.ErrorLabel.Text = uiUtils.EmptyFieldError
        return valid
    }

    if !uiUtils.IsValidSelect(form.PackagingSelect.Select.Selected) {
        valid = false
        form.PackagingSelect.ErrorLabel.Text = uiUtils.EmptyFieldError
        return valid
    }
      
    if !uiUtils.IsValidSelect(form.UnitSElect.Select.Selected) {
        valid = false
        form.UnitSElect.ErrorLabel.Text = uiUtils.EmptyFieldError
        return valid
    }

    return valid
}

