package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
)


func CreatePackagingSelectOptions(w fyne.Window) []string{
	packages, err := services.GetPackagingService()
	if err != nil {
		dialog.NewError(err, w).Show()
	}

	var packNames []string

	for _, pack := range packages {
		packNames = append(packNames, pack.Name)
	}

	return packNames
}


func CreatePackagingSelectWithError(w fyne.Window) *uiTypes.SelectWithError {
	packNames := CreatePackagingSelectOptions(w)
	packSelect := selects.CreateSelectWithError(
        &selects.CreateSelectParams{
			Placeholder: "Выберите упаковку товара",
            Options: packNames,
        })
	return packSelect
}




func GetPackagingID(w fyne.Window, packName string) uint {
	packID, err := services.GetPackagingIDBYNameService(packName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}


	return packID

}