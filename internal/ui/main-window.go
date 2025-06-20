package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {

    w := my_app.NewWindow("Вет склад 2025")

    label := canvas.NewText("Welcome to VetStore2025", color.Black)
    buttonAdd := widget.NewButton("Добавить категорию", func() {
		AddCategoryDialog(w)
	})
    buttonView := widget.NewButton("Показать категории", func() {})

    content := container.NewVBox(label, buttonAdd, buttonView,)
    w.SetContent(content)
    w.Resize(fyne.NewSize(400, 300))
    w.ShowAndRun()
}