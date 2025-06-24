package store_window

import (
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

func ShowStoreWindow() {

	w := ui_utils.CreateNewWindow("ВЕТСКЛАД", false)

	img := ui_utils.CreateWindowImage("static/drugs.png")

	title := "ВЕТСКЛАД"
	

    content := ui_utils.CreateNavWindowContent(img, title, CreateStoreMenu(w),)

    w.SetContent(content)
   
    w.Show()
}