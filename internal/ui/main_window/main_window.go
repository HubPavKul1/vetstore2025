package main_window

import (
	"github.com/HubPavKul1/vetstore2025/internal/ui/ui_utils"
)

func RunUI() {

	w := ui_utils.CreateNewWindow("ВЕТСКЛАД 2025", true)

	img := ui_utils.CreateWindowImage("static/drugs.png")
	menu := CreateMainMenu(w)

	content := ui_utils.CreateNavWindowContent(img, "ДОБРО ПОЖАЛОВАТЬ НА ВЕТСКЛАД", menu)

    w.SetContent(content)
   
    w.Show()
	
}