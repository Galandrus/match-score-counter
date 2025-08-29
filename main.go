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

	// Crear el ViewModel principal
	gameViewModel := viewmodels.NewGameViewModel()

	// Crear la vista principal
	gameView := views.NewGameView(gameViewModel)

	// TODO: Pestaña 2: Configuración
	// configTab := views.NewConfigView(gameViewModel)
	// tabs.Append(container.NewTabItem("Configuración", configTab))

	window.SetContent(gameView.GetMainContainer())
	windowControl.SetContent(gameView.GetControlContainer())

	window.Show()
	windowControl.Show()

	window.ShowAndRun()
	// windowControl.ShowAndRun()
}
