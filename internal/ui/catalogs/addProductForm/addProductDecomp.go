package addProductForm

import (

	"fyne.io/fyne/v2/container"

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

    pack_select, packSelectError := catalogs.CreatePackagingSelectWithError(w)
    unitSelect := catalogs.CreateUnitSelectWithError(w)
    subcatSelect := catalogs.CreateSubCategorySelectWithError()
    cat_select := catalogs.CreateCategorySelectWithError(w)
    cat_select.Select.OnChanged = func(s string) {
        cat_select.ErrorLabel.Text = ""
        if !uiUtils.IsNotEmptyField(s) {
            return
        }
        subcatNames := catalogs.CreateSubCategorySelectOptions(w, s)
        subcatSelect.Select.SetOptions(subcatNames)
    }

    nameEntry := entries.EntryWithError("Введите наименование товара")

    saveButton := uiUtils.CreateSaveBtn()
    backButton := uiUtils.CreateBackBtn(w)

    saveButton.OnTapped = func() {

        // if !IsAddProductFormValid(
            // &AddProductFormField{FieldValue: cat_select.Select.Selected, FieldError: cat_select.ErrorLabel},
            // &AddProductFormField{FieldValue: subcat_select.Selected, FieldError: subcatSelectError},
            // &AddProductFormField{FieldValue: nameEntry.Text, FieldError: nameEntryError},
            // &AddProductFormField{FieldValue: pack_select.Selected, FieldError: packSelectError},
            // &AddProductFormField{FieldValue: unit_select.Selected, FieldError: unitSelectError},
        // ) 
        {
            // return
        }

        selectedSubCategory := subcatSelect.Select.Selected
        subcategoryID := catalogs.GetSubcatID(w, selectedSubCategory)
        name := nameEntry.Input.Text
        selectedPackName := pack_select.Selected
        packID := catalogs.GetPackagingID(w, selectedPackName)
        selectedUnitName := unitSelect.Select.Selected
        unitID := catalogs.GetUnitID(w, selectedUnitName)

        // Создаем новый товар
        SaveNewProduct(w, &AddProductFormData{
            SubcategoryID: subcategoryID,
            PackagingID: packID,
            UnitID: unitID,
            ProductName: name,
        })

        cat_select.Select.ClearSelected()
        subcatSelect.Select.ClearSelected()
        nameEntry.Input.SetText("")
        pack_select.ClearSelected()
        unitSelect.Select.ClearSelected()
    }

    // Обновляем селекты после добавления новых данных
    go HandleUpdateCategoryChannel(updateCategoryChan, w, cat_select.Select)
    go HandleUpdatePackagingChannel(updatePackagingChan, w, pack_select)
    go HandleUpdateUnitChannel(updateUnitChan, w, unitSelect.Select)

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        container.NewHBox(
            container.NewVBox(
                container.NewHBox(cat_select.Select, catalogs.AddCategoryBtn(w, updateCategoryChan)),
                cat_select.ErrorLabel, 
            ),
            container.NewVBox(
                container.NewHBox(subcatSelect.Select, catalogs.AddSubCategoryBtn(w)),
                subcatSelect.ErrorLabel,
            ),   
        ),
        container.NewVBox(nameEntry.Input, nameEntry.ErrorLabel,),
        container.NewHBox(
            container.NewVBox(
                container.NewHBox(pack_select, catalogs.AddPackagingBtn(w, updatePackagingChan)),
                packSelectError,
            ),
            container.NewVBox(
                container.NewHBox(unitSelect.Select, catalogs.AddUnitBtn(w, updateUnitChan)),
                unitSelect.ErrorLabel,
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


