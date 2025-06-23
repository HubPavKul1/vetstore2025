package vets_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

func ShowVetsWindow() {

    w := ui.MyApp.NewWindow("Ветврачи")
	w.Resize(ui_utils.WindowMinSize)
	w.CenterOnScreen()

    label := canvas.NewText("ВЕТВРАЧИ", ui_utils.WindowTitleColor)
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.TextSize = 50
	label.Alignment = fyne.TextAlignCenter



    content := container.NewWithoutLayout(
		label, 
		// buttonAddCat, 
		// buttonAddSubCat, 
		// buttonView,
	)
    w.SetContent(content)
   
    w.Show()
}