package ui_utils

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/HubPavKul1/vetstore2025/internal/ui/app"
)


func CreateColoredButton(c color.RGBA, b *widget.Button, btnText *canvas.Text) *fyne.Container {
	r := canvas.NewRectangle(MenuButtonColor)
	r.StrokeWidth = 1
	r.StrokeColor = c

	btnText.Alignment = fyne.TextAlignCenter

	button := container.New(
		layout.NewStackLayout(),
		r,
		b,
		btnText,
		

	)

	return button

}

func CreateNewWindow(title string, isMaster bool) fyne.Window{
	w := app.MyApp.NewWindow(title)
	w.CenterOnScreen()
	w.Resize(WindowMinSize)
	
	if isMaster {
		w.SetMaster()
	}

	return w
}

func CreateWindowTitle(title string) *canvas.Text {
	text := canvas.NewText(title, WindowTitleColor)
	text.TextSize = 50
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true}
	
	return text
}


func CreateNavWindowContent(img *fyne.Container, title string, menu *fyne.Container) *fyne.Container {
	window_title := CreateWindowTitle(title)
	top_padding := canvas.NewText("", color.Transparent)
	top_padding.TextSize = 10
	
	c := container.NewVBox(
		top_padding,
		container.NewCenter(img),
		window_title,
		widget.NewLabel(""),
		container.NewCenter(menu),
	)

	return c
}

func CreateBtnMenu(btns []*widget.Button) *fyne.Container {

	menu_box := container.NewVBox()

	for _, el := range btns {
		menu_box.Add(el)
	}

	menu_wrapper := container.New(
		layout.NewGridWrapLayout(MenuButtonSize),
		menu_box,
	)

	return menu_wrapper
}


func CreateWindowImage(image_path string) *fyne.Container {
	img := canvas.NewImageFromFile(image_path)
	// img.FillMode = canvas.ImageFillContain

	img_wrapper := container.NewGridWrap(WindowImageSize, img)
	
	
	return img_wrapper
}