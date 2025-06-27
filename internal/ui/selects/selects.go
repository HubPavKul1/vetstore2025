package selects

import (
	"strings"

	"fyne.io/fyne/v2/widget"
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