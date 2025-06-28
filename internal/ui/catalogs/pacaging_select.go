package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
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


func CreatePackagingSelect(w fyne.Window) *widget.Select {
	packNames := CreatePackagingSelectOptions(w)
	pack_select := selects.CreateSelect(
        &selects.CreateSelectParams{
			Placeholder: "Выберите упаковку товара",
            Options: packNames,
        })
	return pack_select
}


func GetPackagingID(w fyne.Window, packName string) uint {
	packID, err := services.GetPackagingIDBYNameService(packName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}


	return packID

}