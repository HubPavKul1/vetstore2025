package addProductForm

import (

	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddProductWindow() {
    // Создаем новое окно
    w := uiUtils.CreateNewWindow("ДОБАВИТЬ ТОВАР В КАТАЛОГ", false)

    //Создаем каналы для обновления селектов после добавления новых данных
    updateCategoryChan := make(chan struct{})
    updatePackagingChan := make(chan struct{})
    updateUnitChan := make(chan struct{})

    packSelect := catalogs.CreatePackagingSelectWithError(w)
    unitSelect := catalogs.CreateUnitSelectWithError(w)
    subcatSelect := catalogs.CreateSubCategorySelectWithError()
    catSelect := catalogs.CreateCategorySelectWithError(w)
    catSelect.Select.OnChanged = func(s string) {
        catSelect.ErrorLabel.Text = ""
        if !uiUtils.IsNotEmptyField(s) {
            return
        }
        subcatNames := catalogs.CreateSubCategorySelectOptions(w, s)
        subcatSelect.Select.SetOptions(subcatNames)
    }

    productNameEntry := entries.EntryWithError("Введите наименование товара")

    saveButton := uiUtils.CreateSaveBtn()
    backButton := uiUtils.CreateBackBtn(w)

    form := AddProductForm{
        CategorySelect: catSelect,
        SubcategorySelect: subcatSelect,
        ProductNameEntry: productNameEntry,
        PackagingSelect: packSelect,
        UnitSelect: unitSelect,
        SaveButton: saveButton,
        BackButton: backButton,
    }

    content := CreateAddProductFormContent(w, &form, updateCategoryChan, updatePackagingChan, updateUnitChan)

    // Обновляем селекты после добавления новых данных
    go HandleUpdateCategoryChannel(updateCategoryChan, w, catSelect.Select)
    go HandleUpdatePackagingChannel(updatePackagingChan, w, packSelect.Select)
    go HandleUpdateUnitChannel(updateUnitChan, w, unitSelect.Select)

    // Устанавливаем контент окна
    w.SetContent(content)

    // Показываем окно
    w.Show()
}