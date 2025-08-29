package main

import (
	"cestoballCounter/src/viewmodels"
	"cestoballCounter/src/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.SetIcon(theme.ColorChromaticIcon())
	window := myApp.NewWindow("Cestoball Score Counter")
	window.Resize(fyne.NewSize(1000, 700))

	windowControl := myApp.NewWindow("Cestoball Score Control")
	window.Resize(fyne.NewSize(1000, 700))

	gameViewModel := viewmodels.NewGameViewModel()
	gameView := views.NewGameView(gameViewModel)

	window.SetContent(gameView.GetMainContainer())
	windowControl.SetContent(gameView.GetControlContainer())

	window.Show()
	windowControl.Show()

	myApp.Run()
}
