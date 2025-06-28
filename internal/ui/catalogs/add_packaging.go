package catalogs

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/dialogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/entries"
)

// AddItemDialog создает диалоговое окно для добавления упаковки
func AddPackagingDialog(parent fyne.Window, updateChan chan<- struct{}) {
    // Создаем новое окно
    dialog_win := dialogs.CreateAddDataDialog(parent, "Добавить упаковку товара")

    // Поле для ввода данных
    name_entry := entries.NameEntry("Введите наименование упаковки")

    form := widget.NewForm(widget.NewFormItem("", name_entry))
    form.SubmitText = "СОХРАНИТЬ"
    form.OnSubmit = func() {
        // Получаем введенные данные
        name := name_entry.Text

        // Создаем новую упаковку
        newPackaging := models.Packaging{}
        newPackaging.Name = name

        // Сохраняем упаковку в базе данных
        _, err := services.CreatePackagingService(newPackaging)
        if err != nil {
            dialog.NewError(err, parent).Show()
            return
        } 
        dialogs.SuccessAddDataDialog(parent).Show() 
        updateChan <- struct{}{}
            
        // Закрываем окно
        dialog_win.Close()
    }

    // Создаем контейнер для полей и кнопки
    content := container.NewVBox(
        form,
    )

    // Устанавливаем контент окна
    dialog_win.SetContent(content)

    // Показываем окно
    dialog_win.Show()
}

func AddPackagingBtn(parent fyne.Window, updateChan chan<- struct{}) *widget.Button {
    btn := widget.NewButton("", func() {AddPackagingDialog(parent, updateChan)})
    btn.Text = strings.ToUpper("Добавить упаковку товара")
    return btn
}