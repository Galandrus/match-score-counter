package src

import (
	"cestoballCounter/src/controllers"
	guimodels "cestoballCounter/src/guiModels"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func CreateMainTab(gameController *controllers.GameController) fyne.CanvasObject {
	mainDisplay := CreateMainDisplay(gameController)
	controlPanel := CreateControlPanel(gameController)
	return container.NewHSplit(mainDisplay, controlPanel)
}

func CreateMainDisplay(gameController *controllers.GameController) fyne.CanvasObject {
	title := guimodels.NewDefaultText("MARCADOR", 40)

	homeTeamGui := gameController.GetHomeTeamGui()
	awayTeamGui := gameController.GetAwayTeamGui()
	clockGui := gameController.GetClockGui()

	homeScoreDisplay := homeTeamGui.TeamContainer
	awayScoreDisplay := awayTeamGui.TeamContainer

	scoresArea := container.NewHBox(homeScoreDisplay, layout.NewSpacer(), awayScoreDisplay)
	mainBox := container.NewVBox(title, scoresArea, layout.NewSpacer(), clockGui.Container, layout.NewSpacer())
	return mainBox
}

func CreateControlPanel(gameController *controllers.GameController) fyne.CanvasObject {
	clockLabel := guimodels.NewDefaultText("CONTROL PANEL", 50)

	homeTeamGui := gameController.GetHomeTeamGui()
	awayTeamGui := gameController.GetAwayTeamGui()
	clockGui := gameController.GetClockGui()

	scoreButtons := container.NewHBox(layout.NewSpacer(), homeTeamGui.ScoreButtons, layout.NewSpacer(), awayTeamGui.ScoreButtons, layout.NewSpacer())

	return container.NewVBox(clockLabel, clockGui.ControlButtons, layout.NewSpacer(), scoreButtons, layout.NewSpacer())
}
