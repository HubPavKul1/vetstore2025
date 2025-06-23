package ui_utils

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


func CreateColoredButton(c color.RGBA, b *widget.Button, btnText *canvas.Text) *fyne.Container {
	r := canvas.NewRectangle(MenuButtonColor)
	r.StrokeWidth = 1
	r.StrokeColor = c

	btnText.Alignment = fyne.TextAlignCenter

	button := container.New(
		layout.NewStackLayout(),
		r,
		b,
		btnText,
		

	)

	return button

}

func CreateWindowTitle(title string, c color.RGBA) *canvas.Text {
	text := canvas.NewText(title, c)
	text.TextSize = 50
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true}
	
	return text
}


func CreateNavWindowContent(img *fyne.Container, title string, menu *fyne.Container) *fyne.Container {
	window_title := CreateWindowTitle(title, WindowTitleColor)
	c := container.NewVBox(
		container.NewCenter(img),
		window_title,
		widget.NewLabel(""),
		container.NewCenter(menu),
	)

	return c
}

func CreateBtnMenu(btns []*widget.Button) *fyne.Container {

	menu_box := container.NewVBox()

	for _, el := range btns {
		menu_box.Add(el)
	}

	menu_wrapper := container.New(
		layout.NewGridWrapLayout(MenuButtonSize),
		menu_box,
	)

	return menu_wrapper
}