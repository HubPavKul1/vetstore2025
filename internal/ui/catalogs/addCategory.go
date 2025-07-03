package catalogs

import (

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


func AddCategoryBtn(parent fyne.Window, updateChan chan<- struct{}) *widget.Button {
    btn := uiUtils.CreateBaseBtn("Добавить категорию товара")
    btn.OnTapped = func() {addCategoryDialog(parent, updateChan)}
    return btn
}


func addCategoryDialog(parent fyne.Window, updateChan chan<- struct{}) {
    // Создаем новое окно
    dialog_win := dialogs.CreateAddDataDialog(parent, "Добавить категорию товара")

    // Поле для ввода данных
    nameEntry := entries.EntryWithError("Введите наименование категории")
    categoryInput := container.NewVBox(nameEntry.Input, nameEntry.ErrorLabel)

    form := widget.NewForm(widget.NewFormItem("", categoryInput))
    form.SubmitText = "СОХРАНИТЬ"
    form.OnSubmit = func() {
        // Получаем введенные данные
        name := nameEntry.Input.Text
        if !uiUtils.IsNotEmptyField(name) {
            nameEntry.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }

        // Создаем новую категорию
        saveNewCategory(parent, updateChan, name)
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


func saveNewCategory(parent fyne.Window, updateChan chan<- struct{}, catName string) {
    // Создаем новую категорию
    newCategory := models.Category{}
    newCategory.Name = catName

    // Сохраняем товар в базе данных
    _, err := services.CreateCategoryService(newCategory)
    if err != nil {
        dialog.NewError(err, parent).Show()
        return
    } 
    dialogs.SuccessAddDataDialog(parent).Show() 
    updateChan <- struct{}{}             

}