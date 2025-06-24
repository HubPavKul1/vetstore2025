package vets_window

import (
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

func ShowVetsWindow() {

    w := ui_utils.CreateNewWindow("ВЕТВРАЧИ", false)

	img := ui_utils.CreateWindowImage("static/treatment.jpeg")

	title := "ВЕТВРАЧИ"
	

    content := ui_utils.CreateNavWindowContent(img, title, CreateVetsMenu(w),)

    w.SetContent(content)
   
    w.Show()
}