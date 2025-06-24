package catalogs

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
)



func CreateSubCategorySelectOptions(w fyne.Window, categoryName string) ([]string, []models.SubCategory ){

	categories, er := services.GetCategoriesService()
	if er != nil {
		dialog.NewError(er, w)
	}
    var categoryID uint
    for _, category := range categories {
    	if category.Name == categoryName {
            categoryID = category.ID
            break
        }
    }

	subcategories, err := services.GetSubCategoriesForCategory(categoryID)
	if err != nil {
		dialog.NewError(err, w)
	}
	var subcats []string
	for _, subcat := range subcategories {
		subcats = append(subcats, subcat.Name)
	}



	return subcats, subcategories

}