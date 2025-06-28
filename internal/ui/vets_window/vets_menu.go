package vets_window

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)


func CreateVetsMenu(window fyne.Window) *fyne.Container{

	saleBtn := widget.NewButton("НАЧАТЬ ПРИЕМ", func() {})

	storeBtn := widget.NewButton("ВЫПИСАТЬ ТОВАР НА ПРИЕМ", func() {})

	backBtn := ui_utils.CreateBackBtn(window)

	btns := []*widget.Button{saleBtn, storeBtn, backBtn, }

	menu := ui_utils.CreateBtnMenu(btns)

	return menu
}