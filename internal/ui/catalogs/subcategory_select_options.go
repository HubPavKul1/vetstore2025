package catalogs

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
)


func CreateSubCategorySelectOptions(w fyne.Window, categoryName string) ([]string, []models.SubCategory ){


	subcategories, err := services.GetSubCategoriesForCategory(categoryName)
	if err != nil {
		dialog.NewError(err, w)
	}
	var subcatNames []string
	for _, subcat := range subcategories {
		subcatNames = append(subcatNames, subcat.Name)
	}

	return subcatNames, subcategories

}