package addProductForm

import (
	"fyne.io/fyne/v2"
	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
)

// SubmitFormHandler определяет процесс сохранения данных формы
func SubmitFormHandler(w fyne.Window, form *AddProductForm) {
	// saveButton := form.SaveButton
    // saveButton.OnTapped = func() {

        // Производим валидацию данных
        if !IsAddProductFormValid(form) {
            return
        }

        // Получаем элементы формы
        categorySelect := form.CategorySelect.Select
        subcategorySelect := form.SubcategorySelect.Select
        productNameInput := form.ProductNameEntry.Input
        packagingSelect := form.PackagingSelect.Select
        unitSelect := form.UnitSelect.Select

        // Получаем IDs из баз данных
        subcategoryID := catalogs.GetSubcatID(w, subcategorySelect.Selected)
        packID := catalogs.GetPackagingID(w, packagingSelect.Selected)
        unitID := catalogs.GetUnitID(w, unitSelect.Selected)
        productName := productNameInput.Text

        // Готовим объект продукта
        product := &AddProductFormData{
            SubcategoryID: subcategoryID,
            ProductName: productName,
            PackagingID:        packID,
            UnitID:        unitID,
        }

        // Сохраняем продукт
        SaveNewProduct(w, product)

        // Очищаем форму
        categorySelect.ClearSelected()
        subcategorySelect.ClearSelected()
        productNameInput.SetText("")
        packagingSelect.ClearSelected()
        unitSelect.ClearSelected()
    }
// }