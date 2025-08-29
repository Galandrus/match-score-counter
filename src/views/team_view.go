package views

import (
	"cestoballCounter/src/viewmodels"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TeamView representa la vista de un equipo con data binding automático
type TeamView struct {
	ViewModel     *viewmodels.TeamViewModel
	TeamLabel     *canvas.Text
	ScoreLabel    *canvas.Text
	TeamContainer *fyne.Container
	ScoreButtons  *fyne.Container
}

// NewTeamView crea una nueva vista de equipo con data binding automático
func NewTeamView(viewModel *viewmodels.TeamViewModel) *TeamView {
	teamView := &TeamView{
		ViewModel: viewModel,
	}

	// Crear elementos de la UI
	teamLabel := canvas.NewText(viewModel.GetName(), nil)
	teamLabel.TextSize = 40
	teamLabel.Alignment = fyne.TextAlignCenter

	scoreLabel := canvas.NewText(strconv.Itoa(viewModel.GetScore()), nil)
	scoreLabel.TextSize = 40
	scoreLabel.Alignment = fyne.TextAlignCenter

	teamView.TeamLabel = teamLabel
	teamView.ScoreLabel = scoreLabel

	// Crear contenedor del equipo
	teamContainer := container.NewVBox(teamLabel, scoreLabel)
	teamView.TeamContainer = teamContainer

	// Crear botones de puntuación
	scoreDoubleButton := widget.NewButton("+2", func() {
		viewModel.AddScore(2)
	})
	scoreTripleButton := widget.NewButton("+3", func() {
		viewModel.AddScore(3)
	})
	scoreMinusButton := widget.NewButton("-1", func() {
		viewModel.AddScore(-1)
	})

	scoreButtons := container.NewVBox(teamLabel, scoreDoubleButton, scoreTripleButton, scoreMinusButton)
	teamView.ScoreButtons = scoreButtons

	// Configurar data binding automático
	teamView.setupDataBinding()

	return teamView
}

// setupDataBinding configura el data binding automático
func (t *TeamView) setupDataBinding() {
	// Binding para el nombre del equipo
	t.ViewModel.Name.Subscribe(func(value interface{}) {
		t.TeamLabel.Text = value.(string)
		t.TeamLabel.Refresh()
	})

	// Binding para el puntaje
	t.ViewModel.Score.Subscribe(func(value interface{}) {
		t.ScoreLabel.Text = strconv.Itoa(value.(int))
		t.ScoreLabel.Refresh()
	})
}

// GetTeamContainer retorna el contenedor del equipo
func (t *TeamView) GetTeamContainer() *fyne.Container {
	return t.TeamContainer
}

// GetScoreButtons retorna los botones de puntuación
func (t *TeamView) GetScoreButtons() *fyne.Container {
	return t.ScoreButtons
}
