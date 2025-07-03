package vetsUI

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)


func CreateVetsMenu(window fyne.Window) *fyne.Container{

	saleBtn := widget.NewButton("НАЧАТЬ ПРИЕМ", func() {})

	storeBtn := widget.NewButton("ВЫПИСАТЬ ТОВАР НА ПРИЕМ", func() {})

	backBtn := uiUtils.CreateBackBtn(window)

	btns := []*widget.Button{saleBtn, storeBtn, backBtn, }

	menu := uiUtils.CreateBtnMenu(btns)

	return menu
}