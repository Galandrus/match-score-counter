package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func CreateMainTab(game *Game) fyne.CanvasObject {
	mainDisplay := CreateMainDisplay(game)
	controlPanel := CreateControlPanel(game)
	return container.NewHSplit(mainDisplay, controlPanel)
}

func CreateMainDisplay(game *Game) fyne.CanvasObject {
	title := NewDefaultText("MARCADOR", 40)

	homeScoreDisplay := game.HomeTeam.TeamContainer
	awayScoreDisplay := game.AwayTeam.TeamContainer

	scoresArea := container.NewHBox(homeScoreDisplay, layout.NewSpacer(), awayScoreDisplay)
	mainBox := container.NewVBox(title, scoresArea, layout.NewSpacer(), game.Clock.Container, layout.NewSpacer())
	return mainBox
}

func CreateControlPanel(game *Game) fyne.CanvasObject {
	clockLabel := NewDefaultText("CONTROL PANEL", 50)

	scoreButtons := container.NewHBox(layout.NewSpacer(), game.HomeTeam.ScoreButtons, layout.NewSpacer(), game.AwayTeam.ScoreButtons, layout.NewSpacer())

	return container.NewVBox(clockLabel, game.Clock.ControlButtons, layout.NewSpacer(), scoreButtons, layout.NewSpacer())
}
