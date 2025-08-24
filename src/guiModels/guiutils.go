package guimodels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewDefaultText(text string, size float32) *canvas.Text {
	textCanvas := canvas.NewText(text, nil)
	textCanvas.TextSize = size
	textCanvas.TextStyle = fyne.TextStyle{Bold: true}
	textCanvas.Alignment = fyne.TextAlignCenter

	return textCanvas
}
