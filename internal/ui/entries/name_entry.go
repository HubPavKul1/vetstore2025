package entries

import (
	"strings"

	"fyne.io/fyne/v2/widget"
)



func NameEntry(placeholder string) *widget.Entry {
	e := widget.NewEntry()
	e.PlaceHolder = strings.ToUpper(placeholder)

	return e
}