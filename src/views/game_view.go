package views

import (
	guimodels "cestoballCounter/src/guiModels"
	"cestoballCounter/src/viewmodels"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// GameView representa la vista principal del juego
type GameView struct {
	ViewModel        *viewmodels.GameViewModel
	HomeTeam         *TeamView
	AwayTeam         *TeamView
	Clock            *ClockView
	RestClock        *RestClockView
	TimeOutClock     *TimeOutClockView
	MainContainer    fyne.CanvasObject
	ControlContainer fyne.CanvasObject
}

// NewGameView crea una nueva vista del juego
func NewGameView(viewModel *viewmodels.GameViewModel) *GameView {
	gameView := &GameView{
		ViewModel: viewModel,
	}

	// Crear vistas de equipos
	gameView.HomeTeam = NewTeamView(viewModel.HomeTeam)
	gameView.AwayTeam = NewTeamView(viewModel.AwayTeam)
	gameView.Clock = NewClockView(viewModel.Clock, viewModel)
	gameView.RestClock = NewRestClockView(viewModel.RestClock)
	gameView.TimeOutClock = NewTimeOutClockView(viewModel.TimeOutClock)

	// Crear contenedor principal
	gameView.MainContainer = gameView.createMainContainer()
	gameView.ControlContainer = gameView.createControlPanel()

	return gameView
}

// createMainContainer crea el contenedor principal de la interfaz
func (g *GameView) createMainContainer() fyne.CanvasObject {
	// Título principal
	title := canvas.NewText("MARCADOR", nil)
	title.TextSize = 40
	title.Alignment = fyne.TextAlignCenter

	// Área de puntuaciones
	homeScoreDisplay := g.HomeTeam.GetTeamContainer()
	awayScoreDisplay := g.AwayTeam.GetTeamContainer()
	scoresArea := container.NewHBox(homeScoreDisplay, layout.NewSpacer(), awayScoreDisplay)

	// Área principal del marcador
	mainDisplay := container.NewVBox(
		title,
		scoresArea,
		layout.NewSpacer(),
		g.Clock.GetContainer(),
		layout.NewSpacer(),
		g.RestClock.GetContainer(),
		layout.NewSpacer(),
		g.TimeOutClock.GetContainer(),
		layout.NewSpacer(),
	)

	return mainDisplay
}

// createControlPanel crea el panel de control
func (g *GameView) createControlPanel() fyne.CanvasObject {
	controlLabel := guimodels.NewDefaultText("CONTROLES", 50)

	// Botones de puntuación
	homeScoreButtons := g.HomeTeam.GetScoreButtons()
	awayScoreButtons := g.AwayTeam.GetScoreButtons()
	scoreButtons := container.NewHBox(
		layout.NewSpacer(),
		homeScoreButtons,
		layout.NewSpacer(),
		awayScoreButtons,
		layout.NewSpacer(),
	)

	// Botones de control del reloj
	clockControlButtons := g.Clock.GetControlButtons()
	restClockControlButtons := g.RestClock.GetControlButtons()
	timeOutClockControlButtons := g.TimeOutClock.GetControlButtons()

	return container.NewVBox(
		controlLabel,
		clockControlButtons,
		layout.NewSpacer(),
		restClockControlButtons,
		layout.NewSpacer(),
		timeOutClockControlButtons,
		layout.NewSpacer(),
		scoreButtons,
		layout.NewSpacer(),
	)
}

// GetContainer retorna el contenedor principal
func (g *GameView) GetMainContainer() fyne.CanvasObject {
	return g.MainContainer
}

// GetControlContainer retorna el contenedor de control
func (g *GameView) GetControlContainer() fyne.CanvasObject {
	return g.ControlContainer
}

// Refresh actualiza toda la vista
func (g *GameView) Refresh() {
	g.MainContainer.Refresh()
	g.ControlContainer.Refresh()
}
