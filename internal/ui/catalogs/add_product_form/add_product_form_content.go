package add_product_form

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
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
    cat_select, catSelectError := catalogs.CreateCategorySelectWithError(w)
    nameEntry, nameEntryError := entries.EntryWithError("Введите наименование товара")
    saveButton := ui_utils.CreateSaveBtn()
    backButton := ui_utils.CreateBackBtn(w)

    content := container.NewVBox(
        container.NewHBox(container.NewVBox(
            container.NewHBox(cat_select, catalogs.AddCategoryBtn(w, updateCategoryChan)), catSelectError,
        )),
        container.NewHBox(container.NewVBox(
            container.NewHBox(subcat_select, catalogs.AddSubCategoryBtn(w)), subcatSelectError,
        )),
        container.NewVBox(nameEntry, nameEntryError),
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
