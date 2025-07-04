package selects

import (
	"strings"

	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiTypes"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)

type CreateSelectParams struct {
	Placeholder string
	Options []string

}

func CreateSelect(params *CreateSelectParams) *widget.Select {
	s := widget.NewSelect([]string{}, func(s string) {})
	s.PlaceHolder = strings.ToUpper(params.Placeholder)
	if len(params.Options) > 0 {
		s.SetOptions(params.Options)
	}

	return s
}


func CreateSelectWithError(params *CreateSelectParams) *uiTypes.SelectWithError {
	s := CreateSelect(params)
	errorLabel := uiUtils.EmptyFieldErrorLabel()
	s.OnChanged = func(s string) {errorLabel.Text = ""}
	selectWithError := uiTypes.SelectWithError{Select: s, ErrorLabel: errorLabel}

	return &selectWithError
}