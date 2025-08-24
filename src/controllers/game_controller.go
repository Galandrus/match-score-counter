package controllers

import (
	"time"

	guimodels "cestoballCounter/src/guiModels"
	"cestoballCounter/src/models"
)

// GameController coordina entre los modelos de negocio y las vistas GUI
type GameController struct {
	homeTeamGui *guimodels.TeamGuiModel
	awayTeamGui *guimodels.TeamGuiModel
	clockGui    *guimodels.ClockGuiModel
	gameConfig  GameConfig
}

// GameConfig contiene la configuración del juego
type GameConfig struct {
	Quarters         int
	TimePerQuarter   time.Duration
	RestTime         time.Duration
	TimePerTimeOut   time.Duration
	TimePerExtraTime time.Duration
}

// DefaultGameConfig configuración por defecto del juego
var DefaultGameConfig = GameConfig{
	Quarters:         2,
	TimePerQuarter:   20 * time.Minute,
	RestTime:         5 * time.Minute,
	TimePerTimeOut:   1 * time.Minute,
	TimePerExtraTime: 3 * time.Minute,
}

// NewGameController crea una nueva instancia del controlador del juego
func NewGameController() *GameController {
	homeTeam := models.Team{Name: "Local", Score: 0}
	awayTeam := models.Team{Name: "Visita", Score: 0}

	homeTeamGui := guimodels.NewTeamGuiModel(homeTeam)
	awayTeamGui := guimodels.NewTeamGuiModel(awayTeam)
	clockGui := guimodels.NewClockGuiModel(DefaultGameConfig.TimePerQuarter)

	return &GameController{
		homeTeamGui: &homeTeamGui,
		awayTeamGui: &awayTeamGui,
		clockGui:    &clockGui,
		gameConfig:  DefaultGameConfig,
	}
}

// GetHomeTeamGui retorna la vista GUI del equipo local
func (gc *GameController) GetHomeTeamGui() *guimodels.TeamGuiModel {
	return gc.homeTeamGui
}

// GetAwayTeamGui retorna la vista GUI del equipo visitante
func (gc *GameController) GetAwayTeamGui() *guimodels.TeamGuiModel {
	return gc.awayTeamGui
}

// GetClockGui retorna la vista GUI del reloj
func (gc *GameController) GetClockGui() *guimodels.ClockGuiModel {
	return gc.clockGui
}

// AddScoreToHomeTeam añade puntos al equipo local
func (gc *GameController) AddScoreToHomeTeam(score int) {
	gc.homeTeamGui.AddScore(score)
}

// AddScoreToAwayTeam añade puntos al equipo visitante
func (gc *GameController) AddScoreToAwayTeam(score int) {
	gc.awayTeamGui.AddScore(score)
}

// GetHomeTeamScore retorna el puntaje del equipo local
func (gc *GameController) GetHomeTeamScore() int {
	return gc.homeTeamGui.GetTeam().GetScore()
}

// GetAwayTeamScore retorna el puntaje del equipo visitante
func (gc *GameController) GetAwayTeamScore() int {
	return gc.awayTeamGui.GetTeam().GetScore()
}

// GetTimeLeft retorna el tiempo restante
func (gc *GameController) GetTimeLeft() time.Duration {
	return gc.clockGui.GetClock().TimeLeft
}
