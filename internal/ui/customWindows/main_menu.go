package customwindows

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	
	"github.com/HubPavKul1/vetstore2025/internal/ui"
)


func MainMenu(window fyne.Window) fyne.Container{

	vetsBtn := widget.NewButton("", func() {ShowVetsWindow()})
	btnText := canvas.NewText("ВЕТВРАЧИ", color.Black)
	coloredVets := ui.CreateColoredButton(ui.MenuButtonColor, vetsBtn, btnText)

	

	storeBtn := widget.NewButton("ВЕТСКЛАД", func() {})

	adminBtn := widget.NewButton("АДМИНИСТРАТОР", func() {})

	closeBtn := widget.NewButton("ВЫЙТИ ИЗ ПРИЛОЖЕНИЯ", func() {window.Close()})

	menu_box := container.NewVBox(
		coloredVets,
		// vetsBtn,
		storeBtn,
		adminBtn,
		closeBtn,
	)

	menu_box_wrapper := container.New(
		layout.NewGridWrapLayout(ui.MenuButtonSize),
		menu_box,
	)

	c := container.New(
		layout.NewCenterLayout(),
		menu_box_wrapper,
	)

	return *c
}