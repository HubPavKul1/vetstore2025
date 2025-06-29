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
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

// Вызывает диалоговое окно для добавления единицы учета
func AddUnitBtn(parent fyne.Window, updateChan chan<- struct{}) *widget.Button {
    btn := widget.NewButton("", func() {addUnitDialog(parent, updateChan)})
    btn.Text = strings.ToUpper("Добавить единицу учета товара")
    return btn
}

func addUnitDialog(parent fyne.Window, updateChan chan<- struct{}) {
    // Создаем новое окно
    dialog_win := dialogs.CreateAddDataDialog(parent, "Добавить единицу учета товара")

    // Поле для ввода данных
    name_entry, errorLabel := entries.EntryWithError("Введите наименование единицы учета")
    unit_input := container.NewVBox(name_entry, errorLabel)

    form := widget.NewForm(widget.NewFormItem("", unit_input))
    form.SubmitText = "СОХРАНИТЬ"
    form.OnSubmit = func() {
        // Получаем введенные данные
        name := name_entry.Text
        if !ui_utils.IsNotEmptyField(name) {
            errorLabel.Text = ui_utils.EmptyFieldError
            return
        }

        // Создаем новую единицу учета
        saveNewUnit(parent, updateChan, name)
        name_entry.SetText("")
       
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

func saveNewUnit(parent fyne.Window, updateChan chan<- struct{}, unitName string) {
    newUnit := models.Unit{}
    newUnit.Name = unitName

    // Сохраняем единицу учета в базе данных
     _, err := services.CreateUnitService(newUnit)
    if err != nil {
        dialog.NewError(err, parent).Show()
        return
    } 
    dialogs.SuccessAddDataDialog(parent).Show() 
    updateChan <- struct{}{}         
}