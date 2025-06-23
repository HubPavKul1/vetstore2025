package main_window

import (

	"github.com/HubPavKul1/vetstore2025/internal/ui"
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

func RunUI() {

    w := ui.MyApp.NewWindow("Вет склад 2025")
	w.Resize(ui_utils.WindowMinSize)
	w.SetMaster()
	
	w.CenterOnScreen()

	img := MainImage(ui_utils.MainImageSize)
	menu := MainMenu(w)

	content := ui_utils.CreateNavWindowContent(img, "ДОБРО ПОЖАЛОВАТЬ НА ВЕТСКЛАД", menu)

    w.SetContent(content)
   
    w.Show()
	
}