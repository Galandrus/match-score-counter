package main

import (
	"cestoballCounter/src/viewmodels"
	"cestoballCounter/src/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.SetIcon(theme.ColorChromaticIcon())
	window := myApp.NewWindow("Cestoball Score Counter - MVVM")
	window.Resize(fyne.NewSize(1000, 700))

	// Crear el ViewModel principal
	gameViewModel := viewmodels.NewGameViewModel()

	// Crear la vista principal
	gameView := views.NewGameView(gameViewModel)

	// Crear pestañas
	tabs := container.NewAppTabs()

	// Pestaña 1: Marcador Principal
	mainScoreTab := gameView.GetContainer()
	tabs.Append(container.NewTabItem("Marcador", mainScoreTab))

	// TODO: Pestaña 2: Configuración
	// configTab := views.NewConfigView(gameViewModel)
	// tabs.Append(container.NewTabItem("Configuración", configTab))

	window.SetContent(tabs)
	window.ShowAndRun()
}
