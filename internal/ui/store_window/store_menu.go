package store_window

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)


func CreateStoreMenu(window fyne.Window) *fyne.Container{

	addProductsBtn := widget.NewButton("ДОБАВИТЬ ТОВАР НА СКЛАД", func() {})

	storeBtn := widget.NewButton("ВЫПИСАТЬ ТОВАР СО СКЛАДА", func() {})

	reportsBtn := widget.NewButton("ОТЧЕТЫ", func() {})

	backBtn := widget.NewButton("НАЗАД", func() {window.Close()})


	btns := []*widget.Button{addProductsBtn, storeBtn, reportsBtn, backBtn, }

	menu := ui_utils.CreateBtnMenu(btns)

	return menu
}