package entries

import (
	"strings"

	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)



func NameEntry(placeholder string) *widget.Entry {
	e := widget.NewEntry()
	e.PlaceHolder = strings.ToUpper(placeholder)

	return e
}


func EntryWithError(placeholder string) *uiTypes.InputWithError {
	e := NameEntry(placeholder)
	errorLabel := uiUtils.EmptyFieldErrorLabel()
	e.OnChanged = func(s string) {errorLabel.Text = ""}
	inputWithError := uiTypes.InputWithError{Input: e, ErrorLabel: errorLabel}

	return &inputWithError
}
