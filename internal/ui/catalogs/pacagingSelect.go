package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
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

func CreatePackagingSelectWithError(w fyne.Window) (*widget.Select, *canvas.Text) {
	pack_select := CreatePackagingSelect(w)
	pack_select_error := uiUtils.EmptyFieldErrorLabel()
	pack_select.OnChanged = func(s string) {pack_select_error.Text = ""}
	return pack_select, pack_select_error
}


func GetPackagingID(w fyne.Window, packName string) uint {
	packID, err := services.GetPackagingIDBYNameService(packName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}


	return packID

}