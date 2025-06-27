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
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
)

// AddItemDialog создает диалоговое окно для добавления товара
func AddSubCategoryDialog(parent fyne.Window) {
    // Создаем новое окно
    dialog_win := dialogs.CreateAddDataDialog(parent, "Добавить подкатегорию")

    // Поле для ввода данных
	categoryNames := CreateCategorySelectOptions(parent)
    categorySelect := selects.CreateSelect(
        &selects.CreateSelectParams{
            Placeholder: "Выберите категорию товара",
            Options: categoryNames,
        },
    )
    
    name_entry := entries.NameEntry("Введите наименование подкатегории товара")

    form := widget.NewForm(
        widget.NewFormItem("", categorySelect),
        widget.NewFormItem("", name_entry),
    )
    form.SubmitText = "СОХРАНИТЬ"
    form.OnSubmit = func() {

        // Получаем введенные данные
        name := name_entry.Text

        selectedCategory := categorySelect.Selected
        categoryID := GetCategoryID(parent, selectedCategory)
        // for _, category := range categories {
        //     if category.Name == selectedCategory {
        //         categoryID = category.ID
        //         break
        //     }
        // }

        // Создаем новую подкатегорию
        newSubCategory := models.SubCategory{}
        newSubCategory.Name = name
        newSubCategory.CategoryID = categoryID

        // Сохраняем подкатегорию в базе данных
        _, err := services.CreateSubCategoryService(newSubCategory)
        if err != nil {
            dialog.NewError(err, parent).Show()
            return
        }

        dialogs.SuccessAddDataDialog(parent).Show()
        
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