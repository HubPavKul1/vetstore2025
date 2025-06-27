package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
)


func CreateUnitSelectOptions(w fyne.Window) []string {
	units, err := services.GetUnitsService()
	if err != nil {
		dialog.NewError(err, w)
	}

	var unitNames []string

	for _, unit := range units {
		unitNames = append(unitNames, unit.Name)
	}

	return unitNames
}

func GetUnitID(w fyne.Window, unitName string) uint {
	unitID, err := services.GetUnitIDBYNameService(unitName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}

	return unitID

}