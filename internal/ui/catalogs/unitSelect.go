package catalogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/HubPavKul1/vetstore2025/internal/services"
	"github.com/HubPavKul1/vetstore2025/internal/ui/selects"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
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


func CreateUnitSelectWithError(w fyne.Window) *uiTypes.SelectWithError {
	unitNames := CreateUnitSelectOptions(w)
	unit_select := selects.CreateSelectWithError(
        &selects.CreateSelectParams{
            Placeholder: "Выберите единицу учета товара",
            Options: unitNames,
        })
	return unit_select
}


// func CreateUnitSelectWithError(w fyne.Window) (*widget.Select, *canvas.Text) {
// 	unit_select := CreateUnitSelect(w)
// 	unit_select_error := uiUtils.EmptyFieldErrorLabel()
// 	unit_select.OnChanged = func(s string) {unit_select_error.Text = ""}
// 	return unit_select, unit_select_error
// }


func GetUnitID(w fyne.Window, unitName string) uint {
	unitID, err := services.GetUnitIDBYNameService(unitName)
	if err != nil {
		dialog.NewError(err, w).Show()
	}

	return unitID

}