package entries

import (
	"strings"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)



func NameEntry(placeholder string) *widget.Entry {
	e := widget.NewEntry()
	e.PlaceHolder = strings.ToUpper(placeholder)

	return e
}


func EntryWithError(placeholder string) (*widget.Entry, *canvas.Text) {
	e := NameEntry(placeholder)
	error_label := ui_utils.EmptyFieldErrorLabel()
	e.OnChanged = func(s string) {error_label.Text = ""}

	return e, error_label
}