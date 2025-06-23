package main_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func MainImage(image_size fyne.Size) *fyne.Container {

	img := canvas.NewImageFromFile("static/vaccinesBg.png")
	img.FillMode = canvas.ImageFillContain
	

	img_wrapper := container.NewGridWrap(image_size, img)
	
	return img_wrapper
}