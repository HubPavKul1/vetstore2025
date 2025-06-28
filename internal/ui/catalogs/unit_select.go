package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
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


func CreateUnitSelect(w fyne.Window) *widget.Select {
	unitNames := CreateUnitSelectOptions(w)
	unit_select := selects.CreateSelect(
        &selects.CreateSelectParams{
            Placeholder: "Выберите единицу учета товара",
            Options: unitNames,
        })
	return unit_select
}


func GetUnitID(w fyne.Window, unitName string) uint {
	unitID, err := services.GetUnitIDBYNameService(unitName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}

	return unitID

}