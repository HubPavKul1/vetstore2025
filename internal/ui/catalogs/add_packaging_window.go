package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddPackagingDialog(parent fyne.Window) {
    
    w := dialogs.CreateAddDataDialog(parent, "Добавить упаковку товара")

    nameEntry := widget.NewEntry()

    // Кнопка для подтверждения
    saveButton := widget.NewButton("Сохранить", func() {
        // Получаем введенные данные
        name := nameEntry.Text

        // Создаем новую упаковку
        newPackaging := models.Packaging{}
        newPackaging.Name = name

        // Сохраняем товар в базе данных
        _, err := services.CreatePackagingService(newPackaging)
        if err != nil {
            dialog.NewError(err, parent).Show()
            return
        } 
        dialogs.SuccessAddDataDialog(parent).Show()
            
        // Закрываем окно
        w.Close()


        // Обновляем список товаров в главном окне
        // (здесь нужно реализовать логику обновления списка)
    })

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        nameEntry,
        saveButton,
    )

    // Устанавливаем контент окна
    w.SetContent(content)

    // Показываем окно
    w.Show()
}