package ui_types

import (

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)


type InputWithError struct {
	Input *widget.Entry
	ErrorLabel *canvas.Text
}

type SelectWithError struct {
	Select *widget.Select
	ErrorLabel *canvas.Text
}
