package add_product_form

import (

	"fyne.io/fyne/v2/container"

	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddProductWindow() {
    // Создаем новое окно
    w := ui_utils.CreateNewWindow("ДОБАВИТЬ ТОВАР В КАТАЛОГ", false)

    //Создаем каналы для обновления селектов после добавления новых данных
    updateCategoryChan := make(chan struct{})
    updatePackagingChan := make(chan struct{})
    updateUnitChan := make(chan struct{})

    pack_select, packSelectError := catalogs.CreatePackagingSelectWithError(w)
    unit_select, unitSelectError := catalogs.CreateUnitSelectWithError(w)
    subcat_select, subcatSelectError := catalogs.CreateSubCategorySelectWithError()
    cat_select := catalogs.CreateCategorySelectWithError(w)
    cat_select.Select.OnChanged = func(s string) {
        cat_select.ErrorLabel.Text = ""
        if !ui_utils.IsNotEmptyField(s) {
            return
        }
        subcatNames := catalogs.CreateSubCategorySelectOptions(w, s)
        subcat_select.SetOptions(subcatNames)
    }

    nameEntry, nameEntryError := entries.EntryWithError("Введите наименование товара")

    saveButton := ui_utils.CreateSaveBtn()
    backButton := ui_utils.CreateBackBtn(w)

    saveButton.OnTapped = func() {

        if !IsAddProductFormValid(
            &AddProductFormField{FieldValue: cat_select.Select.Selected, FieldError: cat_select.ErrorLabel},
            &AddProductFormField{FieldValue: subcat_select.Selected, FieldError: subcatSelectError},
            &AddProductFormField{FieldValue: nameEntry.Text, FieldError: nameEntryError},
            &AddProductFormField{FieldValue: pack_select.Selected, FieldError: packSelectError},
            &AddProductFormField{FieldValue: unit_select.Selected, FieldError: unitSelectError},
        ) {
            return
        }

        selectedSubCategory := subcat_select.Selected
        subcategoryID := catalogs.GetSubcatID(w, selectedSubCategory)
        name := nameEntry.Text
        selectedPackName := pack_select.Selected
        packID := catalogs.GetPackagingID(w, selectedPackName)
        selectedUnitName := unit_select.Selected
        unitID := catalogs.GetUnitID(w, selectedUnitName)

        // Создаем новый товар
        SaveNewProduct(w, &AddProductFormData{
            SubcategoryID: subcategoryID,
            PackagingID: packID,
            UnitID: unitID,
            ProductName: name,
        })

        cat_select.Select.ClearSelected()
        subcat_select.ClearSelected()
        nameEntry.SetText("")
        pack_select.ClearSelected()
        unit_select.ClearSelected()
    }

    // Обновляем селекты после добавления новых данных
    go HandleUpdateCategoryChannel(updateCategoryChan, w, cat_select.Select)
    go HandleUpdatePackagingChannel(updatePackagingChan, w, pack_select)
    go HandleUpdateUnitChannel(updateUnitChan, w, unit_select)

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        container.NewHBox(
            container.NewVBox(
                container.NewHBox(cat_select.Select, catalogs.AddCategoryBtn(w, updateCategoryChan)),
                cat_select.ErrorLabel, 
            ),
            container.NewVBox(
                container.NewHBox(subcat_select, catalogs.AddSubCategoryBtn(w)),
                subcatSelectError,
            ),   
        ),
        container.NewVBox(nameEntry, nameEntryError,),
        container.NewHBox(
            container.NewVBox(
                container.NewHBox(pack_select, catalogs.AddPackagingBtn(w, updatePackagingChan)),
                packSelectError,
            ),
            container.NewVBox(
                container.NewHBox(unit_select, catalogs.AddUnitBtn(w, updateUnitChan)),
                unitSelectError,
            ),    
        ),
        saveButton,
        backButton,
    )

    // Устанавливаем контент окна
    w.SetContent(content)

    // Показываем окно
    w.Show()
}


