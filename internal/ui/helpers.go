package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


func CreateColoredButton(c color.RGBA, b *widget.Button, btnText *canvas.Text) *fyne.Container {
	r := canvas.NewRectangle(c)
	r.StrokeWidth = 0

	button := container.New(
		layout.NewStackLayout(),
		b,
		r,
		btnText,
		

	)

	return button

}