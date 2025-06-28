package catalogs

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddProduct() {
    // Создаем новое окно
    w := ui_utils.CreateNewWindow("ДОБАВИТЬ ТОВАР В КАТАЛОГ", false)

    //Создаем каналы для обновления селектов после добавления новых данных
    updateCategoryChan := make(chan struct{})
    updatePackagingChan := make(chan struct{})
    updateUnitChan := make(chan struct{})

    pack_select, packSelectError := CreatePackagingSelectWithError(w)
    unit_select, unitSelectError := CreateUnitSelectWithError(w)
    subcat_select, subcatSelectError := CreateSubCategorySelectWithError()
    cat_select := CreateCategorySelect(w)
    cat_select.OnChanged = func(s string) {
        subcatNames := CreateSubCategorySelectOptions(w, s)
        subcat_select.SetOptions(subcatNames)
    }

    nameEntry, nameEntryError := entries.EntryWithError("Введите наименование товара")

    saveButton := ui_utils.CreateSaveBtn()
    backButton := ui_utils.CreateBackBtn(w)

    // Кнопка для подтверждения
    saveButton.OnTapped = func() {
    	// Получаем введенные данные
        valid := true

        // Получаем выбранную подкатегорию
        selectedSubCategory := subcat_select.Selected
        if !ui_utils.IsValidSelect(selectedSubCategory) {
            valid = false
            subcatSelectError.Text = ui_utils.EmptyFieldError
            return
        }
        log.Println("SElected SubCat: ", selectedSubCategory)
        subcategoryID := GetSubcatID(w, selectedSubCategory)
        log.Println("SubcatID for New Product: ", subcategoryID)
        
        // Получаем наименование товара
        name := nameEntry.Text
        if !ui_utils.IsNotEmptyField(name) {
            valid = false
            nameEntryError.Text = ui_utils.EmptyFieldError
            return
        }

        // Получаем ID упаковки
        selectedPackName := pack_select.Selected
        if !ui_utils.IsValidSelect(selectedPackName) {
            valid = false
            packSelectError.Text = ui_utils.EmptyFieldError
            return
        }
        // packSelectError.Text = ""
        packID := GetPackagingID(w, selectedPackName)
        log.Println("UNITID IS: ", packID)
    
        // Получаем ID единицы учета
        selectedUnitName := unit_select.Selected
        if !ui_utils.IsValidSelect(selectedUnitName) {
            valid = false
            unitSelectError.Text = ui_utils.EmptyFieldError
            return
        }

        unitID := GetUnitID(w, selectedUnitName)
        log.Println("UNITID IS: ", unitID)

        if !valid {
            log.Println("Форма НеВалидна!!!")
            return
        }

        log.Println("ISVALID FORM: ", valid)
        //     // Создаем новый товар
        newProduct := models.Product{
            SubCategoryID: subcategoryID, 
            PackagingID: packID,
            UnitID: unitID,
        }
        newProduct.Name = name
        // Сохраняем товар в базе данных
        _, err := services.CreateProductService(newProduct)
            if err != nil {
                dialog.NewError(err, w).Show()
                return
            }

        dialogs.SuccessAddDataDialog(w).Show()
		
    }

    // Обновляем селекты после добавления новых данных
    go HandleUpdateCategoryChannel(updateCategoryChan, w, cat_select)
    go HandleUpdatePackagingChannel(updatePackagingChan, w, pack_select)
    go HandleUpdateUnitChannel(updateUnitChan, w, unit_select)

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        container.NewHBox(cat_select, AddCategoryBtn(w, updateCategoryChan)),
        container.NewVBox(
            container.NewHBox(subcat_select, AddSubCategoryBtn(w)),subcatSelectError,
        ),
        container.NewVBox(nameEntry, nameEntryError,),
        container.NewVBox(
            container.NewHBox(pack_select, AddPackagingBtn(w, updatePackagingChan)),
            packSelectError,
        ),
        container.NewVBox(
            container.NewHBox(unit_select, AddUnitBtn(w, updateUnitChan)),
            unitSelectError,
        ),
        saveButton,
        backButton,
    )

    // Устанавливаем контент окна
    w.SetContent(content)

    // Показываем окно
    w.Show()
}