package addProductForm

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
)


func CreateAddProductFormContent(
	w fyne.Window, 
	form *AddProductForm,
	updateCategoryChan chan<- struct{},
	updatePackagingChan chan<- struct{},
	updateUnitChan chan<- struct{},
	) *fyne.Container {


    content := container.NewVBox(
        container.NewHBox(container.NewVBox(
            container.NewHBox(form.CategorySelect.Select, catalogs.AddCategoryBtn(w, updateCategoryChan)), form.CategorySelect.ErrorLabel,
        )),
        container.NewHBox(container.NewVBox(
            container.NewHBox(form.SubcategorySelect.Select, catalogs.AddSubCategoryBtn(w)), form.SubcategorySelect.ErrorLabel,
        )),
        container.NewVBox(form.ProductNameEntry.Input, form.ProductNameEntry.ErrorLabel),
        container.NewHBox(container.NewVBox(
            container.NewHBox(form.PackagingSelect.Select, catalogs.AddPackagingBtn(w, updatePackagingChan)), form.PackagingSelect.ErrorLabel,
        )),
        container.NewHBox(container.NewVBox(
            container.NewHBox(form.UnitSelect.Select, catalogs.AddUnitBtn(w, updateUnitChan)), form.UnitSelect.ErrorLabel,
        )),
        form.SaveButton,
        form.BackButton,
    )

    form.SaveButton.OnTapped = func() {SubmitFormHandler(w, form)}

    return content
}
