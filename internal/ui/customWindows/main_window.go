package customwindows

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/HubPavKul1/vetstore2025/internal/ui"
)

func RunUI() {

    w := ui.MyApp.NewWindow("Вет склад 2025")
	w.SetMaster()
	w.Resize(ui.WindowSize)
	w.CenterOnScreen()

	// r := canvas.NewRectangle(color.NRGBA{255, 255, 255, 255})

	// window_container := container.NewCenter(r)

	img := canvas.NewImageFromFile("static/vaccinesBg.png")
	img.FillMode = canvas.ImageFillContain
	img_x_size := ui.WindowSize.Width
	img_y_size := ui.WindowSize.Height / 2
	img_wrapper := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(img_x_size, img_y_size)),
		img,
	)

	c := container.New(
		layout.NewCenterLayout(),
		img_wrapper,
	)

    title := canvas.NewText("ДОБРО ПОЖАЛОВАТЬ НА ВЕТСКЛАД", ui.WindowTitleColor)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 50
	title.Alignment = fyne.TextAlignCenter

	menu := MainMenu(w)


    content := container.NewVBox(
		c,
		title, 
		widget.NewLabel(""),
		&menu,
	)
    w.SetContent(content)
   
    w.Show()
	
}