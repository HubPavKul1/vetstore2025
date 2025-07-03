package addProductForm

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
)

func HandleUpdateCategoryChannel(
    updateCategoryChan <-chan struct{}, 
    w fyne.Window, 
    category_select *widget.Select,
) {
    for range updateCategoryChan {
        catNames := catalogs.CreateCategorySelectOptions(w)
        fyne.Do(func() {category_select.SetOptions(catNames) })
    }
}

func HandleUpdatePackagingChannel(
    updatePackagingChan <-chan struct{}, 
    w fyne.Window, 
    packaging_select *widget.Select,
) {
    for range updatePackagingChan {
        packNames := catalogs.CreatePackagingSelectOptions(w)
        fyne.Do(func() {packaging_select.SetOptions(packNames)})
    }
    
}

func HandleUpdateUnitChannel(
    updateUnitChan <-chan struct{}, 
    w fyne.Window, 
    unit_select *widget.Select,
) {
    for range updateUnitChan {
        unitNames := catalogs.CreateUnitSelectOptions(w)
        fyne.Do(func() {unit_select.SetOptions(unitNames)})
    }
    
}
