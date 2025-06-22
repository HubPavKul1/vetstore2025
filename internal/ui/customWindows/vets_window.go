package customwindows

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"


	// "fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui"
)

func ShowVetsWindow() {

    w := ui.MyApp.NewWindow("Ветврачи")
	w.Resize(ui.WindowSize)
	w.CenterOnScreen()

    label := canvas.NewText("ВЕТВРАЧИ", ui.WindowTitleColor)
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.TextSize = 50
	label.Alignment = fyne.TextAlignCenter

	
    // buttonAddCat := widget.NewButton("Добавить категорию", func() {
	// 	AddCategoryDialog(w)
	// })
	// buttonAddSubCat := widget.NewButton("Добавить подкатегорию", func() {
	// 	AddSubCategoryDialog(w)
	// })
    // buttonView := widget.NewButton("Показать категории", func() {})

    content := container.NewWithoutLayout(
		label, 
		// buttonAddCat, 
		// buttonAddSubCat, 
		// buttonView,
	)
    w.SetContent(content)
   
    w.Show()
}