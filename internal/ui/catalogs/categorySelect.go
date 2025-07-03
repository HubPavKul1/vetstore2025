package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)


func CreateCategorySelectOptions(w fyne.Window) []string {
	categories, err := services.GetCategoriesService()
	if err != nil {
		dialog.NewError(err, w).Show()
	}

	var catNames []string
	for _, cat := range categories {
		catNames = append(catNames, cat.Name)
	}

	return catNames
}

func CreateCategorySelect(w fyne.Window) *widget.Select {
	catOptions := CreateCategorySelectOptions(w)
	catSelect := selects.CreateSelect(&selects.CreateSelectParams{
		Placeholder: "Выберите категорию товара",
		Options: catOptions,
	})
	return catSelect
}

func CreateCategorySelectWithError(w fyne.Window) *uiTypes.SelectWithError {
	cat_select := CreateCategorySelect(w)
	cat_select_error := uiUtils.EmptyFieldErrorLabel()
	cat_select.OnChanged = func(s string) {cat_select_error.Text = ""}

	return &uiTypes.SelectWithError{Select: cat_select, ErrorLabel: cat_select_error}
}

func GetCategoryID(w fyne.Window, catName string) uint {
	catID, err := services.GetCategoryIDBYNameService(catName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}

	return catID

}
