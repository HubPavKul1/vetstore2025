package mainWindow

import (
	"github.com/HubPavKul1/vetstore2025/internal/ui/uiUtils"
)

func RunUI() {

	w := uiUtils.CreateNewWindow("ВЕТСКЛАД 2025", true)

	img := uiUtils.CreateWindowImage("static/drugs.png")
	menu := CreateMainMenu(w)

	content := uiUtils.CreateNavWindowContent(img, "ДОБРО ПОЖАЛОВАТЬ НА ВЕТСКЛАД", menu)

    w.SetContent(content)
   
    w.Show()
	
}