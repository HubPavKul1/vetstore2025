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

    pack_select := CreatePackagingSelect(w)
    unit_select := CreateUnitSelect(w)
    subcat_select := CreateSubCategorySelect()
    cat_select := CreateCategorySelect(w)
    cat_select.OnChanged = func(s string) {
        subcatNames := CreateSubCategorySelectOptions(w, s)
        subcat_select.SetOptions(subcatNames)
    }

    nameEntry := entries.NameEntry("Введите наименование товара")

    saveButton := ui_utils.CreateSaveBtn()
    backButton := ui_utils.CreateBackBtn(w)

    // Кнопка для подтверждения
    saveButton.OnTapped = func() {
    	// Получаем введенные данные
        name := nameEntry.Text
        // Получаем выбранную подкатегорию
        selectedSubCategory := subcat_select.Selected
        log.Println("SElected SubCat: ", selectedSubCategory)
        subcategoryID := GetSubcatID(w, selectedSubCategory)
        log.Println("SubcatID for New Product: ", subcategoryID)

        // Получаем ID упаковки
        selectedPackName := pack_select.Selected
        packID := GetPackagingID(w, selectedPackName)
    
        // Получаем ID единицы учета
        selectedUnitName := unit_select.Selected
        unitID := GetUnitID(w, selectedUnitName)
        log.Println("UNITID IS: ", unitID)

		// Создаем новый товар
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
        container.NewHBox(subcat_select, AddSubCategoryBtn(w)),
        nameEntry,
        container.NewHBox(pack_select, AddPackagingBtn(w, updatePackagingChan)),
        container.NewHBox(unit_select, AddUnitBtn(w, updateUnitChan)),
        saveButton,
        backButton,
    )

    // Устанавливаем контент окна
    w.SetContent(content)

    // Показываем окно
    w.Show()
}
