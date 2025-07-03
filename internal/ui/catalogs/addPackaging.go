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
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)

// AddItemDialog вызывает диалоговое окно для добавления упаковки
func AddPackagingBtn(parent fyne.Window, updateChan chan<- struct{}) *widget.Button {
    btn := widget.NewButton("", func() {addPackagingDialog(parent, updateChan)})
    btn.Text = strings.ToUpper("Добавить упаковку товара")
    return btn
}

func addPackagingDialog(parent fyne.Window, updateChan chan<- struct{}) {
    // Создаем новое окно
    dialog_win := dialogs.CreateAddDataDialog(parent, "Добавить упаковку товара")

    // Поле для ввода данных
    nameEntry := entries.EntryWithError("Введите наименование упаковки")
    packInput := container.NewVBox(nameEntry.Input, nameEntry.ErrorLabel)

    form := widget.NewForm(widget.NewFormItem("", packInput))
    form.SubmitText = "СОХРАНИТЬ"
    form.OnSubmit = func() {
        // Получаем введенные данные
        name := nameEntry.Input.Text
        if !uiUtils.IsNotEmptyField(name) {
            nameEntry.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }

        // Сохраняем в базу данных
        saveNewPackaging(parent, updateChan, name)
        nameEntry.Input.SetText("")
            
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

func saveNewPackaging(parent fyne.Window, updateChan chan<- struct{}, packName string) {
     // Создаем новую упаковку
    newPackaging := models.Packaging{}
    newPackaging.Name = packName

    // Сохраняем упаковку в базе данных
    _, err := services.CreatePackagingService(newPackaging)
    if err != nil {
        dialog.NewError(err, parent).Show()
        return
    } 
    dialogs.SuccessAddDataDialog(parent).Show() 
    updateChan <- struct{}{}
}