package customwindows

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"


	// "fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui"
)

func RunUI() {

    w := ui.MyApp.NewWindow("Вет склад 2025")
	w.SetMaster()
	w.Resize(fyne.NewSize(1500, 800))
	img := canvas.NewImageFromFile("static/vaccinesBg.png")
	
	img.Resize(fyne.NewSize(800, 400))
	img.Move(fyne.NewPos(400, 30))

    label := canvas.NewText("ДОБРО ПОЖАЛОВАТЬ НА ВЕТСКЛАД", color.Black)
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.TextSize = 50
	label.Move(fyne.NewPos(400, 500))

	
    // buttonAddCat := widget.NewButton("Добавить категорию", func() {
	// 	AddCategoryDialog(w)
	// })
	// buttonAddSubCat := widget.NewButton("Добавить подкатегорию", func() {
	// 	AddSubCategoryDialog(w)
	// })
    // buttonView := widget.NewButton("Показать категории", func() {})

    content := container.NewWithoutLayout(
		img, 
		label, 
		// buttonAddCat, 
		// buttonAddSubCat, 
		// buttonView,
	)
    w.SetContent(content)
   
    w.ShowAndRun()
}