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
        fyne.Do(func() { category_select.SetOptions(catNames) })
    }
}

func HandlUpdatePackagingChannel(
    updatePackagingChan <-chan struct{}, 
    w fyne.Window, 
    packaging_select *widget.Select,
) {
    packNames := CreatePackagingSelectOptions(w)
    fyne.Do(func() {packaging_select.SetOptions(packNames)})
}


func HandlUpdateUnitChannel(
    updateUnitChan <-chan struct{}, 
    w fyne.Window, 
    unit_select *widget.Select,
) {
    unitNames := CreateUnitSelectOptions(w)
    fyne.Do(func() {unit_select.SetOptions(unitNames)})
}
