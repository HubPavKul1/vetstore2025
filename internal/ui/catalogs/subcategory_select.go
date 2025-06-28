package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)


func CreateSubCategorySelectOptions(w fyne.Window, categoryName string) []string {

	subcategories, err := services.GetSubCategoriesForCategoryService(categoryName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}
	var subcatNames []string
	for _, subcat := range subcategories {
		subcatNames = append(subcatNames, subcat.Name)
	}

	return subcatNames

}

func CreateSubCategorySelect() *widget.Select {
	subcat_select := selects.CreateSelect(
		&selects.CreateSelectParams{
			Placeholder: "Выберите подкатегорию товара",
		})
	return subcat_select
}

func CreateSubCategorySelectWithError() (*widget.Select, *canvas.Text) {
	subcat_select := CreateSubCategorySelect()
	subcat_select_error := ui_utils.EmptyFieldErrorLabel()
	subcat_select.OnChanged = func(s string) {subcat_select_error.Text = ""}
	
	return subcat_select, subcat_select_error
}


func GetSubcatID(w fyne.Window, subcatName string) uint {
	subcatID, err := services.GetSubCategoryIDBYNameService(subcatName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}
	
	return subcatID
}