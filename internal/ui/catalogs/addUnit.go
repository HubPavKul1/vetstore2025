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
    nameInput := entries.EntryWithError("Введите наименование единицы учета")
    unitInput := container.NewVBox(nameInput.Input, nameInput.ErrorLabel)

    form := widget.NewForm(widget.NewFormItem("", unitInput))
    form.SubmitText = "СОХРАНИТЬ"
    form.OnSubmit = func() {
        // Получаем введенные данные
        name := nameInput.Input.Text
        if !uiUtils.IsNotEmptyField(name) {
            nameInput.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }

        // Создаем новую единицу учета
        saveNewUnit(parent, updateChan, name)
        nameInput.Input.SetText("")
       
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