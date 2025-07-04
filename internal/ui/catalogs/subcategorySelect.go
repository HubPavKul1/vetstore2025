package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
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

func CreateSubCategorySelectWithError() *uiTypes.SelectWithError {
	subcatSelect := selects.CreateSelectWithError(
		&selects.CreateSelectParams{
			Placeholder: "Выберите подкатегорию товара",
		})
	return subcatSelect
}


func GetSubcatID(w fyne.Window, subcatName string) uint {
	subcatID, err := services.GetSubCategoryIDBYNameService(subcatName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}
	
	return subcatID
}