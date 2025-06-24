package store_window


import (

	"github.com/HubPavKul1/vetstore2025/internal/ui"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

func ShowStoreWindow() {

    w := ui.MyApp.NewWindow("ВЕТСКЛАД")
	w.Resize(ui_utils.WindowMinSize)
	w.CenterOnScreen()

	img := ui_utils.CreateWindowImage("static/drugs.png")

	title := "ВЕТСКЛАД"
	

    content := ui_utils.CreateNavWindowContent(img, title, CreateStoreMenu(w),)

    w.SetContent(content)
   
    w.Show()
}