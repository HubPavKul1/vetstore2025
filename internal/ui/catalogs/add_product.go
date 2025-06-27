package catalogs

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddProduct() {
    // Создаем новое окно
    w := ui_utils.CreateNewWindow("ДОБАВИТЬ ТОВАР В КАТАЛОГ", false)

    updateCategoryChan := make(chan struct{})
    updatePackagingChan := make(chan struct{})
    updateUnitChan := make(chan struct{})

    catNames := CreateCategorySelectOptions(w)
    packNames := CreatePackagingSelectOptions(w)
    unitNames := CreateUnitSelectOptions(w)

    pack_select := selects.CreateSelect(
        &selects.CreateSelectParams{
            Placeholder: "Выберите упаковку товара",
            Options: packNames,
        }, 
    )
    unit_select := selects.CreateSelect(
        &selects.CreateSelectParams{
            Placeholder: "Выберите единицу учета товара",
            Options: unitNames,
        },
         
    )

    subcat_select := selects.CreateSelect(
        &selects.CreateSelectParams{
            Placeholder: "Выберите подкатегорию товара",
        },     
    )

    cat_select := selects.CreateSelect(
        &selects.CreateSelectParams{
            Placeholder: "Выберите категорию товара",
            Options: catNames,
        },     
    )

    cat_select.OnChanged = func(s string) {
        subcatNames := CreateSubCategorySelectOptions(w, s)
        subcat_select.SetOptions(subcatNames)
    }

    addCategoryBtn := widget.NewButton(
        "ДОБАВИТЬ КАТЕГОРИЮ ТОВАРА", 
        func() {AddCategoryDialog(w, updateCategoryChan)},
    )

    addSubcatBtn := widget.NewButton(
        "ДОБАВИТЬ ПОДКАТЕГОРИЮ ТОВАРА", 
        func() {AddSubCategoryDialog(w)},
    )

    addPackagingBtn := widget.NewButton(
        "ДОБАВИТЬ УПАКОВКУ ТОВАРА", 
        func() {AddPackagingDialog(w, updatePackagingChan)},
    )

    addUnitBtn := widget.NewButton(
        "ДОБАВИТЬ ЕДИНИЦУ УЧЕТА ТОВАРА", 
        func() {AddUnitDialog(w, updateUnitChan)},
    )

    nameEntry := entries.NameEntry("Введите наименование товара")

    // Кнопка для подтверждения
    saveButton := widget.NewButton("СОХРАНИТЬ", func() {
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


        // Обновляем список товаров в главном окне
        // (здесь нужно реализовать логику обновления списка)
    })

    go HandleUpdateCategoryChannel(updateCategoryChan, w, cat_select)
    go HandlUpdatePackagingChannel(updatePackagingChan, w, pack_select)
    go HandlUpdateUnitChannel(updateUnitChan, w, unit_select)

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        container.NewHBox(cat_select, addCategoryBtn),
        container.NewHBox(subcat_select, addSubcatBtn),
        nameEntry,
        container.NewHBox(pack_select, addPackagingBtn),
        container.NewHBox(unit_select, addUnitBtn),
        saveButton,
    )

    // Устанавливаем контент окна
    w.SetContent(content)

    // Показываем окно
    w.Show()
}



