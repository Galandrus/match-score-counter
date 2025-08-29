package guimodels

import (
	"cestoballCounter/src/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TeamGuiModel representa la interfaz visual de un equipo
type TeamGuiModel struct {
	team          *models.Team // Referencia al modelo de negocio
	TeamLabel     *canvas.Text
	ScoreLabel    *canvas.Text
	TeamContainer *fyne.Container
	ScoreButtons  *fyne.Container
}

// AddScore añade puntos al equipo y actualiza la interfaz
func (t *TeamGuiModel) AddScore(score int) {
	t.team.AddScore(score)
	t.updateDisplay()
}

// ResetScore reinicia el puntaje del equipo y actualiza la interfaz
func (t *TeamGuiModel) ResetScore() {
	t.team.ResetScore()
	t.updateDisplay()
}

// updateDisplay actualiza la visualización del puntaje
func (t *TeamGuiModel) updateDisplay() {
	t.ScoreLabel.Text = strconv.Itoa(t.team.GetScore())
	t.ScoreLabel.Refresh()
}

// GetTeam retorna el modelo de negocio del equipo
func (t *TeamGuiModel) GetTeam() *models.Team {
	return t.team
}

// NewTeamGuiModel crea una nueva instancia de la interfaz visual del equipo
func NewTeamGuiModel(team models.Team) TeamGuiModel {
	teamGuiModel := TeamGuiModel{team: &team}
	teamLabel := NewDefaultText(team.GetName(), 40)
	scoreLabel := NewDefaultText(strconv.Itoa(team.GetScore()), 40)
	teamContainer := container.NewVBox(teamLabel, scoreLabel)

	teamGuiModel.ScoreLabel = scoreLabel
	teamGuiModel.TeamLabel = teamLabel
	teamGuiModel.TeamContainer = teamContainer

	scoreDoubleButton := widget.NewButton("+2", func() {
		teamGuiModel.AddScore(2)
	})
	scoreTripleButton := widget.NewButton("+3", func() {
		teamGuiModel.AddScore(3)
	})
	scoreMinusButton := widget.NewButton("-1", func() {
		teamGuiModel.AddScore(-1)
	})
	scoreButtons := container.NewVBox(teamLabel, scoreDoubleButton, scoreTripleButton, scoreMinusButton)
	teamGuiModel.ScoreButtons = scoreButtons

	return teamGuiModel
}
