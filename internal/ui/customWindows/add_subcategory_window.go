package customwindows

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/db"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/db/repository"
	service "github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddSubCategoryDialog(parent fyne.Window) {
    // Создаем новое окно
    dialog := ui.MyApp.NewWindow("Добавить подкатегорию")

	categories, err := service.GetCategoriesService()
	if err != nil {
		log.Panic("Ошибка получения категорий")
		return
	} 
	var cats []string
	for _, cat := range categories {
		cats = append(cats, cat.Name)
	}

	
	

    // Поле для ввода данных
	categorySelect := widget.NewSelect([]string{}, func(s string) {})
    categorySelect.PlaceHolder = "Выберите категорию товара"
	categorySelect.SetOptions(cats)
    nameEntry := widget.NewEntry()
	
   

    // Кнопка для подтверждения
    saveButton := widget.NewButton("Сохранить", func() {
    	// Получаем введенные данные
        name := nameEntry.Text
        
        // Получаем выбранную категорию
        selectedCategory := categorySelect.Selected
        var categoryID uint
        for _, category := range categories {
            if category.Name == selectedCategory {
                categoryID = category.ID
                break
            }
        }
		// Создаем новую подкатегорию
        newSubCategory := models.SubCategory{Name: name, CategoryID: categoryID}

        // Сохраняем товар в базе данных
        _, err := repository.CreateSubCategory(db.DB, newSubCategory)
        if err != nil {
            log.Panic(err, parent)
            return
        }

	



        // Закрываем окно
        dialog.Close()

        // Обновляем список товаров в главном окне
        // (здесь нужно реализовать логику обновления списка)
    })

    // Создаем контейнер для полей и кнопки
     content := container.NewVBox(
		widget.NewLabel("Категория:"),
        categorySelect,
        widget.NewLabel("Название:"),
        nameEntry,
        saveButton,
    )

    // Устанавливаем контент окна
    dialog.SetContent(content)

    // Показываем окно
    dialog.Show()
}