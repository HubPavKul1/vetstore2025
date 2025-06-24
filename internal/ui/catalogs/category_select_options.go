package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
)


func CreateCategorySelectOptions(w fyne.Window) ([]string, []models.Category ){
	categories, err := services.GetCategoriesService()
	if err != nil {
		dialog.NewError(err, w)
	}

	var catNames []string

	for _, cat := range categories {
		catNames = append(catNames, cat.Name)
	}

	return catNames, categories

}