package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
)


func CreatePackagingSelect(w fyne.Window) (*widget.Select, []models.Packaging){
	packages, err := services.GetPackagingService()
	if err != nil {
		dialog.NewError(err, w)
	}

	var packs []string

	for _, pack := range packages {
		packs = append(packs, pack.Name)
	}

	packs_select := widget.NewSelect(packs, func(s string) {})
	packs_select.PlaceHolder = "Выберите упаковку"

	return packs_select, packages
}