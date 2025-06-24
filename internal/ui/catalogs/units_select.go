package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/db/models"
	"github.com/HubPavKul1/vetstore2025/internal/services"
)


func CreateUnitsSelect(w fyne.Window) (*widget.Select, []models.Unit){
	units, err := services.GetUnitsService()
	if err != nil {
		dialog.NewError(err, w)
	}

	var unitNames []string

	for _, unit := range units {
		unitNames = append(unitNames, unit.Name)
	}

	units_select := widget.NewSelect(unitNames, func(s string) {})
	units_select.PlaceHolder = "Выберите упаковку"

	return units_select, units
}