package main_window

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui/store_window"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
	"github.com/HubPavKul1/vetstore2025/internal/ui/vets_window"
)


func CreateMainMenu(window fyne.Window) *fyne.Container{

	vetsBtn := widget.NewButton("ВЕТВРАЧИ", func() {vets_window.ShowVetsWindow()})

	storeBtn := widget.NewButton("ВЕТСКЛАД", func() {store_window.ShowStoreWindow()})

	adminBtn := widget.NewButton("АДМИНИСТРАТОР", func() {})

	closeBtn := widget.NewButton("ВЫЙТИ ИЗ ПРИЛОЖЕНИЯ", func() {window.Close()})

	btns := []*widget.Button{vetsBtn, storeBtn, adminBtn, closeBtn,}

	menu := ui_utils.CreateBtnMenu(btns)

	return menu
}