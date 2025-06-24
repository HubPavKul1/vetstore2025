package catalogs

import (
	"log"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/app"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddSubCategoryDialog(parent fyne.Window) {
    // Создаем новое окно
    dialog_win := app.MyApp.NewWindow("Добавить подкатегорию")

    // Поле для ввода данных
	categoryNames, categories:= CreateCategorySelectOptions(parent)
    categorySelect := widget.NewSelect(categoryNames, func(s string) {})
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

        // Сохраняем подкатегорию в базе данных
        _, err := services.CreateSubCategoryService(newSubCategory)
        if err != nil {
            log.Panic(err, parent)
            return
        }

        dialog.ShowInformation("", "Подкатегория успешно создана!", parent,)

        // Закрываем окно
        dialog_win.Close()

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
    dialog_win.SetContent(content)

    // Показываем окно
    dialog_win.Show()
}