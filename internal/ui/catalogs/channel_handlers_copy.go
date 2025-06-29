package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func HandleUpdateCategoryChannel(
    updateCategoryChan <-chan struct{}, 
    w fyne.Window, 
    category_select *widget.Select,
) {
    for range updateCategoryChan {
        catNames := CreateCategorySelectOptions(w)
        fyne.Do(func() {category_select.SetOptions(catNames) })
    }
}

func HandleUpdatePackagingChannel(
    updatePackagingChan <-chan struct{}, 
    w fyne.Window, 
    packaging_select *widget.Select,
) {
    for range updatePackagingChan {
        packNames := CreatePackagingSelectOptions(w)
        fyne.Do(func() {packaging_select.SetOptions(packNames)})
    }
    
}

func HandleUpdateUnitChannel(
    updateUnitChan <-chan struct{}, 
    w fyne.Window, 
    unit_select *widget.Select,
) {
    for range updateUnitChan {
        unitNames := CreateUnitSelectOptions(w)
        fyne.Do(func() {unit_select.SetOptions(unitNames)})
    }
    
}
