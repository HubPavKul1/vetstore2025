package catalogs

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	// "github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddProduct() {
    // Создаем новое окно
    w := ui_utils.CreateNewWindow("ДОБАВИТЬ ТОВАР В КАТАЛОГ", false)

    updateCategoryChan := make(chan bool)

    catNames, _ := CreateCategorySelectOptions(w)

    var subcategories []models.SubCategory

    subcat_select := widget.NewSelect([]string{}, func(s string) {})
    subcat_select.PlaceHolder = "ВЫБЕРИТЕ ПОДКАТЕГОРИЮ ТОВАРА"

    cat_select := widget.NewSelect(catNames, func(s string) {
        
        subcatNames, subcats := CreateSubCategorySelectOptions(w, s)
        subcategories = subcats
        subcat_select.SetOptions(subcatNames)

    })
    cat_select.PlaceHolder = "ВЫБЕРИТЕ КАТЕГОРИЮ ТОВАРА"

    addCategoryBtn := widget.NewButton("ДОБАВИТЬ КАТЕГОРИЮ ТОВАРА", func() {AddCategoryDialog(w, updateCategoryChan)})

    
    addSubcatBtn := widget.NewButton("ДОБАВИТЬ ПОДКАТЕГОРИЮ ТОВАРА", func() {AddSubCategoryDialog(w)})

    nameEntry := widget.NewEntry()
    nameEntry.PlaceHolder = "Введите наименование товара"

    // pack_select, packs := CreatePackagingSelect(w)

    // unit_select, units := CreateUnitsSelect(w)


    // Кнопка для подтверждения
    saveButton := widget.NewButton("Сохранить", func() {
    	// Получаем введенные данные
        // name := nameEntry.Text
        
        // Получаем выбранную подкатегорию
        selectedSubCategory := subcat_select.Selected
        log.Println("SElected SubCat: ", selectedSubCategory)
        var subcategoryID uint
        
        for _, subcat := range subcategories {
            if subcat.Name == selectedSubCategory {
                subcategoryID = subcat.ID
                break
            }
        }
        log.Println("SubcatID for New Product: ", subcategoryID)

        //Получаем ID упаковки
        // selectedPackName := pack_select.Selected
        // var packID uint
        // for _, pack := range packs {
        //     if pack.Name == selectedPackName {
        //         packID = pack.ID
        //         break
        //     }
        // }

        //Получаем ID единицы учета
        // selectedUnitName := unit_select.Selected
        // var unitID uint
        // for _, unit := range units {
        //     if unit.Name == selectedUnitName {
        //         unitID = unit.ID
        //         break
        //     }
        // }

		// Создаем новый товар
        // newProduct := models.Product{
        //     Name: name, 
        //     SubCategoryID: subcategoryID, 
        //     PackagingID: packID,
        //     UnitID: unitID,
        // }

        // Сохраняем товар в базе данных
        // _, err := services.CreateProductService(newProduct)
        // if err != nil {
        //     log.Panic(err)
        //     return
        // }

        // Закрываем окно


        // Обновляем список товаров в главном окне
        // (здесь нужно реализовать логику обновления списка)
    })

    // Создаем контейнер для полей и кнопки
     content := container.NewVBox(
		// widget.NewLabel("Категория:"),
        container.NewHBox(
            cat_select,
            addCategoryBtn,
        ),
        container.NewHBox(
            subcat_select,
            addSubcatBtn,

        ),
    
        nameEntry,
        // pack_select,
        // unit_select,
        saveButton,
    )

    // Устанавливаем контент окна
    w.SetContent(content)

    go func() {
        for range updateCategoryChan {
            catNames, _ := CreateCategorySelectOptions(w)
            fyne.Do(func() {cat_select.SetOptions(catNames)})
           
        }
    }()

    // Показываем окно
    w.Show()
}