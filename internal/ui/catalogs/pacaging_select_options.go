package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
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


func GetPackagingID(w fyne.Window, pacName string) uint {
	packID, err := services.GetPackagingIDBYNameService(pacName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}


	return packID

}