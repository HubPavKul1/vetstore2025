package storeUI

import (
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)

func ShowStoreWindow() {

	w := uiUtils.CreateNewWindow("ВЕТСКЛАД", false)

	img := uiUtils.CreateWindowImage("static/drugs.png")

	title := "ВЕТСКЛАД"
	

    content := uiUtils.CreateNavWindowContent(img, title, CreateStoreMenu(w),)

    w.SetContent(content)
   
    w.Show()
}