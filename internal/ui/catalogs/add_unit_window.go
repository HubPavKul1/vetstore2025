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
func AddUnitsDialog(parent fyne.Window) {
    
    w := dialogs.CreateAddDataDialog(parent, "Добавить единицу учета товара")

    nameEntry := widget.NewEntry()
    nameEntry.PlaceHolder = "ВВЕДИТЕ НАИМЕНОВАНИЕ ЕДИНИЦЫ УЧЕТА ТОВАРА"

    // Кнопка для подтверждения
    saveButton := widget.NewButton("СОХРАНИТЬ", func() {
        // Получаем введенные данные
        name := nameEntry.Text

        // Создаем новую упаковку
        newUnit := models.Unit{}
        newUnit.Name = name

        // Сохраняем товар в базе данных
        _, err := services.CreateUnitService(newUnit)
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