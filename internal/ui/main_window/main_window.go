package main_window

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui"
)

func RunUI() {

    w := ui.MyApp.NewWindow("Вет склад 2025")
	w.SetMaster()
	
	w.CenterOnScreen()

	img := MainImage(ui.MainImageSize)

	title := ui.CreateWindowTitle("ДОБРО ПОЖАЛОВАТЬ НА ВЕТСКЛАД", ui.WindowTitleColor)

	menu := MainMenu(w)

	content := container.NewVBox(
		container.NewCenter(img),
		title,
		widget.NewLabel(""),
		container.NewCenter(menu),

	)


    w.SetContent(content)
   
    w.Show()
	
}