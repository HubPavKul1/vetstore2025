package mainWindow

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui/storeUI"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
	"github.com/HubPavKul1/vetstore2025/internal/ui/vetsUI"
)


func CreateMainMenu(window fyne.Window) *fyne.Container{

	vetsBtn := widget.NewButton("ВЕТВРАЧИ", func() {vetsUI.ShowVetsWindow()})

	storeBtn := widget.NewButton("ВЕТСКЛАД", func() {storeUI.ShowStoreWindow()})

	adminBtn := widget.NewButton("АДМИНИСТРАТОР", func() {})

	closeBtn := widget.NewButton("ВЫЙТИ ИЗ ПРИЛОЖЕНИЯ", func() {window.Close()})

	btns := []*widget.Button{vetsBtn, storeBtn, adminBtn, closeBtn,}

	menu := uiUtils.CreateBtnMenu(btns)

	return menu
}