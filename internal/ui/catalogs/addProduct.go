package catalogs

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddProduct() {
    // Создаем новое окно
    w := uiUtils.CreateNewWindow("ДОБАВИТЬ ТОВАР В КАТАЛОГ", false)

    //Создаем каналы для обновления селектов после добавления новых данных
    updateCategoryChan := make(chan struct{})
    updatePackagingChan := make(chan struct{})
    updateUnitChan := make(chan struct{})

    pack_select, packSelectError := CreatePackagingSelectWithError(w)
    unitSelect := CreateUnitSelectWithError(w)
    subcatSelect := CreateSubCategorySelectWithError()
    cat_select := CreateCategorySelectWithError(w)
    cat_select.Select.OnChanged = func(s string) {
        cat_select.ErrorLabel.Text = ""
        if !uiUtils.IsNotEmptyField(s) {
            return
        }
        subcatNames := CreateSubCategorySelectOptions(w, s)
        subcatSelect.Select.SetOptions(subcatNames)
    }

    nameEntry := entries.EntryWithError("Введите наименование товара")

    saveButton := uiUtils.CreateSaveBtn()
    backButton := uiUtils.CreateBackBtn(w)

    saveButton.OnTapped = func() {

    	// Получаем введенные данные и валидируем их
        valid := true

        selectedCategory := cat_select.Select.Selected
        if !uiUtils.IsValidSelect(selectedCategory) {
            valid = false
            cat_select.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }
        
        // Получаем выбранную подкатегорию
        selectedSubCategory := subcatSelect.Select.Selected
        if !uiUtils.IsValidSelect(selectedSubCategory) {
            valid = false
            subcatSelect.ErrorLabel.Text = uiUtils.EmptyFieldError
            cat_select.Select.ClearSelected()
            return
        }
        subcategoryID := GetSubcatID(w, selectedSubCategory)
        
        // Получаем наименование товара
        name := nameEntry.Input.Text
        if !uiUtils.IsNotEmptyField(name) {
            valid = false
            nameEntry.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }

        // Получаем ID упаковки
        selectedPackName := pack_select.Selected
        if !uiUtils.IsValidSelect(selectedPackName) {
            valid = false
            packSelectError.Text = uiUtils.EmptyFieldError
            return
        }
        // packSelectError.Text = ""
        packID := GetPackagingID(w, selectedPackName)
    
        // Получаем ID единицы учета
        selectedUnitName := unitSelect.Select.Selected
        if !uiUtils.IsValidSelect(selectedUnitName) {
            valid = false
            unitSelect.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }

        unitID := GetUnitID(w, selectedUnitName)

        if !valid {
            return
        }

        // Создаем новый товар
        saveNewProduct(w, &addProductForm{
            SubcategoryID: subcategoryID,
            PackID: packID,
            UnitID: unitID,
            Name: name,
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
                container.NewHBox(cat_select.Select, AddCategoryBtn(w, updateCategoryChan)),
                cat_select.ErrorLabel, 
            ),
            container.NewVBox(
                container.NewHBox(subcatSelect.Select, AddSubCategoryBtn(w)),
                subcatSelect.ErrorLabel,
            ),   
        ),
        container.NewVBox(nameEntry.Input, nameEntry.ErrorLabel,),
        container.NewHBox(
            container.NewVBox(
                container.NewHBox(pack_select, AddPackagingBtn(w, updatePackagingChan)),
                packSelectError,
            ),
            container.NewVBox(
                container.NewHBox(unitSelect.Select, AddUnitBtn(w, updateUnitChan)),
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


type addProductForm struct {
    SubcategoryID uint
    Name string
    PackID uint
    UnitID uint
}

func saveNewProduct( w fyne.Window, f *addProductForm) {
    newProduct := models.Product{
        SubCategoryID: f.SubcategoryID, 
        PackagingID: f.PackID,
        UnitID: f.UnitID,
    }
    newProduct.Name =f.Name
    // Сохраняем товар в базе данных
    _, err := services.CreateProductService(newProduct)
        if err != nil {
            dialog.NewError(err, w).Show()
            return
        }
    dialogs.SuccessAddDataDialog(w).Show()
}


// type formField struct {
//     fieldValue string
//     errorLabel *canvas.Text
// }

// func validator(formFields []*formField) bool {
//     valid := true
//     for _, elem := range formFields {
//         if !ui_utils.IsNotEmptyField(elem.fieldValue) {
//             valid = false
//             elem.errorLabel.Text = ui_utils.EmptyFieldError
//             break
//         }
//         return valid
//     }
//     return valid     
// }