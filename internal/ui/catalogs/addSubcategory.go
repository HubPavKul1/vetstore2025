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

func AddSubCategoryBtn(parent fyne.Window) *widget.Button {
      btn := widget.NewButton("", func() {addSubCategoryDialog(parent)})
      btn.Text = "ДОБАВИТЬ ПОДКАТЕГОРИЮ ТОВАРА"
      return btn
}

func addSubCategoryDialog(parent fyne.Window) {
    // Создаем новое окно
    dialog_win := dialogs.CreateAddDataDialog(parent, "Добавить подкатегорию")

    // Поле для ввода данных
    categorySelect := CreateCategorySelectWithError(parent)
    categorySelectError := container.NewVBox(categorySelect.Select, categorySelect.ErrorLabel)
    
    nameEntry := entries.EntryWithError("Введите наименование подкатегории товара")
    subcategoryInput := container.NewVBox(nameEntry.Input, nameEntry.ErrorLabel)

    form := widget.NewForm(
        widget.NewFormItem("", categorySelectError),
        widget.NewFormItem("", subcategoryInput),
    )
    form.SubmitText = "СОХРАНИТЬ"
    form.OnSubmit = func() {
        valid := true
        // Получаем введенные данные
        selectedCategory := categorySelect.Select.Selected
        if !uiUtils.IsValidSelect(selectedCategory) {
            valid = false
            categorySelect.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }
        categoryID := GetCategoryID(parent, selectedCategory)

        name := nameEntry.Input.Text
        if !uiUtils.IsNotEmptyField(name) {
            valid = false
            nameEntry.ErrorLabel.Text = uiUtils.EmptyFieldError
            return
        }
        
        if !valid {
            return
        }

        // Создаем новую подкатегорию
        saveNewSubcategory(parent, &subcategoryForm{CategoryID: categoryID, Name: name})
        categorySelect.Select.ClearSelected()
        nameEntry.Input.SetText("")
        
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

type subcategoryForm struct {
    CategoryID uint
    Name string
}

func saveNewSubcategory(parent fyne.Window, f *subcategoryForm) {
    newSubCategory := models.SubCategory{}
    newSubCategory.Name = f.Name
    newSubCategory.CategoryID = f.CategoryID

    // Сохраняем подкатегорию в базе данных
    _, err := services.CreateSubCategoryService(newSubCategory)
    if err != nil {
        dialog.NewError(err, parent).Show()
        return
    }

    dialogs.SuccessAddDataDialog(parent).Show()
}