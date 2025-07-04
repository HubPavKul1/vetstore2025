package storeUI

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs"
	"github.com/HubPavKul1/vetstore2025/internal/ui/catalogs/addProductForm"
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)


func CreateStoreMenu(window fyne.Window) *fyne.Container{

	addProductsBtn := widget.NewButton("ДОБАВИТЬ ТОВАР НА СКЛАД", func() {})

	moveOutProdBtn := widget.NewButton("ВЫПИСАТЬ ТОВАР СО СКЛАДА", func() {})

	prodCatalogBtn := widget.NewButton("ОТКРЫТЬ КАТАЛОГ ТОВАРОВ", func() {catalogs.VeiwProductCatalog()})

	// addProductToCatalogBtn := widget.NewButton("ДОБАВИТЬ ТОВАР В КАТАЛОГ", func() {catalogs.AddProduct()})

	addProductToCatalogBtn := widget.NewButton("ДОБАВИТЬ ТОВАР В КАТАЛОГ", func() {addProductForm.AddProductWindow()})

	reportsBtn := widget.NewButton("ОТЧЕТЫ", func() {})

	backBtn := uiUtils.CreateBackBtn(window)


	btns := []*widget.Button{
		addProductsBtn, 
		moveOutProdBtn, 
		prodCatalogBtn, 
		addProductToCatalogBtn,
		reportsBtn, 
		backBtn, 
	}

	menu := uiUtils.CreateBtnMenu(btns)

	return menu
}