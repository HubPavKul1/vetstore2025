package vetsUI

import (
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)

func ShowVetsWindow() {

    w := uiUtils.CreateNewWindow("ВЕТВРАЧИ", false)

	img := uiUtils.CreateWindowImage("static/treatment.jpeg")

	title := "ВЕТВРАЧИ"
	

    content := uiUtils.CreateNavWindowContent(img, title, CreateVetsMenu(w),)

    w.SetContent(content)
   
    w.Show()
}