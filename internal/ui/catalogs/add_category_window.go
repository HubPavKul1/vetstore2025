package catalogs

import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddCategoryDialog(parent fyne.Window, updateChan chan<- bool) {
    // Создаем новое окно
    dialog_win := dialogs.CreateAddDataDialog(parent, "Добавить категорию товара")

    // Поле для ввода данных
    nameEntry := widget.NewEntry()
   

    // Кнопка для подтверждения
    saveButton := widget.NewButton("Сохранить", func() {
        // Получаем введенные данные
        name := nameEntry.Text

        // Создаем новую категорию
        newCategory := models.Category{}
        newCategory.Name = name

        // Сохраняем товар в базе данных
        _, err := services.CreateCategoryService(newCategory)
        if err != nil {
            dialog.NewError(err, parent).Show()
            return
        } 
        dialogs.SuccessAddDataDialog(parent).Show() 
        updateChan <- true
            
        // Закрываем окно
        dialog_win.Close()


        // Обновляем список товаров в главном окне
        // (здесь нужно реализовать логику обновления списка)
    })

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        nameEntry,
        saveButton,
    )

    // Устанавливаем контент окна
    dialog_win.SetContent(content)

    // Показываем окно
    dialog_win.Show()
}