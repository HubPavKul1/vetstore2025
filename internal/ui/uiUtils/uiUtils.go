package uiUtils

import (
	"image/color"
	"strings"

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

func CreateBaseBtn(title string) *widget.Button {
	btn := widget.NewButton("", func() {})
	btn.Text = strings.ToUpper(title)
	return btn
}

func CreateBackBtn(w fyne.Window) *widget.Button {
	btn := CreateBaseBtn("Назад")
	btn.OnTapped = func() {w.Close()}
	return btn
}

func CreateSaveBtn() *widget.Button {
	return CreateBaseBtn("Сохранить")
}


func CreateNewWindow(title string, isMaster bool) fyne.Window{
	w := app.MyApp.NewWindow(strings.ToUpper(title))
	w.CenterOnScreen()
	w.Resize(WindowMinSize)
	
	if isMaster {
		w.SetMaster()
	}

	return w
}

func CreateWindowTitle(title string) *canvas.Text {
	text := canvas.NewText(strings.ToUpper(title), WindowTitleColor)
	text.TextSize = 50
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true}
	
	return text
}


func CreateNavWindowContent(img *fyne.Container, title string, menu *fyne.Container) *fyne.Container {
	windowTitle := CreateWindowTitle(strings.ToUpper(title))
	topPadding := canvas.NewText("", color.Transparent)
	topPadding.TextSize = 10
	
	c := container.NewVBox(
		topPadding,
		container.NewCenter(img),
		windowTitle,
		widget.NewLabel(""),
		container.NewCenter(menu),
	)

	return c
}

func CreateBtnMenu(btns []*widget.Button) *fyne.Container {

	menuBox := container.NewVBox()

	for _, el := range btns {
		menuBox.Add(el)
	}

	menuWrapper := container.New(
		layout.NewGridWrapLayout(MenuButtonSize),
		menuBox,
	)

	return menuWrapper
}


func CreateWindowImage(image_path string) *fyne.Container {
	img := canvas.NewImageFromFile(image_path)
	// img.FillMode = canvas.ImageFillContain

	imgWrapper := container.NewGridWrap(WindowImageSize, img)
	
	
	return imgWrapper
}


