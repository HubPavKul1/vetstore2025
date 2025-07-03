package uiUtils

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2/canvas"

)

func IsValidString(str string, minLength int) bool {
    trimmedStr := strings.TrimSpace(str)
    return len(trimmedStr) >= minLength && trimmedStr != ""
}

func IsNotEmptyField(str string) bool {
	trimmedStr := strings.TrimSpace(str)
    return trimmedStr != ""
}

func IsValidSelect(selected string) bool {
	return selected != ""
}

func EmptyFieldErrorLabel() *canvas.Text {
    errorText := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	return errorText
}

var EmptyFieldError = "Поле не должно быть пустым"