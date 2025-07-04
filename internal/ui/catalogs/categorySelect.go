package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
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

func CreateCategorySelectWithError(w fyne.Window) *uiTypes.SelectWithError {
	catOptions := CreateCategorySelectOptions(w)
	catSelect := selects.CreateSelectWithError(&selects.CreateSelectParams{
		Placeholder: "Выберите категорию товара",
		Options: catOptions,
	})
	return catSelect
}


func GetCategoryID(w fyne.Window, catName string) uint {
	catID, err := services.GetCategoryIDBYNameService(catName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}

	return catID

}
