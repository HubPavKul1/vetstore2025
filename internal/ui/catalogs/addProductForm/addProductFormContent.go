package addProductForm

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)


func CreateAddProductForm(
	w fyne.Window, 
	saveCallback func(),
	updateCategoryChan chan<- struct{},
	updatePackagingChan chan<- struct{},
	updateUnitChan chan<- struct{},
	) *fyne.Container {

    pack_select, packSelectError := catalogs.CreatePackagingSelectWithError(w)
    unit_select, unitSelectError := catalogs.CreateUnitSelectWithError(w)
    subcat_select, subcatSelectError := catalogs.CreateSubCategorySelectWithError()
    cat_select := catalogs.CreateCategorySelectWithError(w)
    nameEntry := entries.EntryWithError("Введите наименование товара")
    saveButton := uiUtils.CreateSaveBtn()
    backButton := uiUtils.CreateBackBtn(w)

    content := container.NewVBox(
        container.NewHBox(container.NewVBox(
            container.NewHBox(cat_select.Select, catalogs.AddCategoryBtn(w, updateCategoryChan)), cat_select.ErrorLabel,
        )),
        container.NewHBox(container.NewVBox(
            container.NewHBox(subcat_select, catalogs.AddSubCategoryBtn(w)), subcatSelectError,
        )),
        container.NewVBox(nameEntry.Input, nameEntry.ErrorLabel),
        container.NewHBox(container.NewVBox(
            container.NewHBox(pack_select, catalogs.AddPackagingBtn(w, updatePackagingChan)), packSelectError,
        )),
        container.NewHBox(container.NewVBox(
            container.NewHBox(unit_select, catalogs.AddUnitBtn(w, updateUnitChan)), unitSelectError,
        )),
        saveButton,
        backButton,
    )

    saveButton.OnTapped = saveCallback

    return content
}
