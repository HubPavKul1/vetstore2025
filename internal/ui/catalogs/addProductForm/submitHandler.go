package addProductForm

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// SubmitFormHandler определяет процесс сохранения данных формы
func SubmitFormHandler(w fyne.Window, form *fyne.Container, formFields []*string) {
	saveButton := form.Objects[len(form.Objects)-2].(*widget.Button)
    saveButton.OnTapped = func() {
        // Сбор данных из формы

        // name := nameEntry.Text
        // category := cat_select.Selected
        // subcategory := subcat_select.Selected
        // packaging := pack_select.Selected
        // unit := unit_select.Selected

        // Производим валидацию данных
        // if !ValidateForm(name, category, subcategory, packaging, unit) {
        //     return
        // }

        // Получаем IDs из баз данных
        // subcategoryID := GetSubcatID(w, subcategory)
        // packID := GetPackagingID(w, packaging)
        // unitID := GetUnitID(w, unit)

        // Готовим объект продукта
        // product := &addProductForm{
        //     SubcategoryID: subcategoryID,
        //     Name:          name,
        //     PackID:        packID,
        //     UnitID:        unitID,
        // }

        // Сохраняем продукт
        // saveNewProduct(w, product)

        // Очищаем форму
        // clearFormFields()
    }
}

// clearFormFields сбрасывает состояние полей формы
// func clearFormFields() {
//     cat_select.ClearSelected()
//     subcat_select.ClearSelected()
//     nameEntry.SetText("")
//     pack_select.ClearSelected()
//     unit_select.ClearSelected()
// }