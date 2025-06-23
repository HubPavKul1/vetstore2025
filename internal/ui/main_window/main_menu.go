package main_window

import (
	

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui"
	"github.com/HubPavKul1/vetstore2025/internal/ui/vets_window"
)


func MainMenu(window fyne.Window) *fyne.Container{

	vetsBtn := widget.NewButton("ВЕТВРАЧИ", func() {vets_window.ShowVetsWindow()})

	storeBtn := widget.NewButton("ВЕТСКЛАД", func() {})

	adminBtn := widget.NewButton("АДМИНИСТРАТОР", func() {})

	closeBtn := widget.NewButton("ВЫЙТИ ИЗ ПРИЛОЖЕНИЯ", func() {window.Close()})


	menu_box := container.NewVBox(
		vetsBtn,
		storeBtn,
		adminBtn,
		closeBtn,
	)

	menu_wrapper := container.New(
		layout.NewGridWrapLayout(ui.MenuButtonSize),
		menu_box,
	)

	return menu_wrapper
}