package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/db"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/db/repository"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddCategoryDialog(parent fyne.Window) {
    // Создаем новое окно
    dialog := my_app.NewWindow("Добавить категорию")

    // Поля для ввода данных
    nameEntry := widget.NewEntry()
   

    // Кнопка для подтверждения
    saveButton := widget.NewButton("Сохранить", func() {
        // Получаем введенные данные
        name := nameEntry.Text

        // Создаем новую категорию
        newCategory := models.Category{Name: name,}

        // Сохраняем товар в базе данных
        _, err := repository.CreateCategory(db.DB, newCategory)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Закрываем окно
        dialog.Close()

        // Обновляем список товаров в главном окне
        // (здесь нужно реализовать логику обновления списка)
    })

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        widget.NewLabel("Название:"),
        nameEntry,
        saveButton,
    )

    // Устанавливаем контент окна
    dialog.SetContent(content)

    // Показываем окно
    dialog.Show()
}