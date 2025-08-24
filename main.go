package main

import (
	"cestoballCounter/src"
	"cestoballCounter/src/controllers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.SetIcon(theme.ColorChromaticIcon())
	window := myApp.NewWindow("Cestoball Score Counter")
	window.Resize(fyne.NewSize(1000, 700))

	gameController := controllers.NewGameController()

	// Crear pestañas
	tabs := container.NewAppTabs()

	// Pestaña 1: Marcador Principal
	mainScoreTab := src.CreateMainTab(gameController)
	tabs.Append(container.NewTabItem("Marcador", mainScoreTab))

	// // Pestaña 2: Configuración
	// configTab := src.CreateConfigTab(gameController)
	// tabs.Append(container.NewTabItem("Configuración", configTab))

	window.SetContent(tabs)
	window.ShowAndRun()
}
