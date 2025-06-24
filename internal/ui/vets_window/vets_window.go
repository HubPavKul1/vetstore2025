package vets_window

import (

	"github.com/HubPavKul1/vetstore2025/internal/ui"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

func ShowVetsWindow() {

    w := ui.MyApp.NewWindow("ВЕТВРАЧИ")
	w.Resize(ui_utils.WindowMinSize)
	w.CenterOnScreen()

	img := ui_utils.CreateWindowImage("static/treatment.jpeg")

	title := "ВЕТВРАЧИ"
	

    content := ui_utils.CreateNavWindowContent(img, title, CreateVetsMenu(w),)

    w.SetContent(content)
   
    w.Show()
}