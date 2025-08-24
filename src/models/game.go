package models

import (
	"time"
)

type GameConfig struct {
	Quarters         int
	TimePerQuarter   time.Duration
	RestTime         time.Duration
	QuarterTime      time.Duration
	TimePerTimeOut   time.Duration
	TimePerExtraTime time.Duration
}

var DefaultGameConfig = GameConfig{
	Quarters:         2,
	TimePerQuarter:   20 * time.Minute,
	RestTime:         5 * time.Minute,
	TimePerTimeOut:   1 * time.Minute,
	TimePerExtraTime: 3 * time.Minute,
}

// Game representa el juego con su lógica de negocio pura
type Game struct {
	HomeTeam Team
	AwayTeam Team
	Config   GameConfig
	Clock    Clock
}

// Start inicia el juego
func (g *Game) Start() {
	g.Clock.Start()
}

// Stop detiene el juego
func (g *Game) Stop() {
	g.Clock.Stop()
}

// Reset reinicia el juego
func (g *Game) Reset() {
	g.HomeTeam.ResetScore()
	g.AwayTeam.ResetScore()
	g.Clock.Reset()
}

// NewGame crea una nueva instancia del juego
func NewGame() *Game {
	return &Game{
		HomeTeam: Team{Name: "Local", Score: 0},
		AwayTeam: Team{Name: "Visita", Score: 0},
		Config:   DefaultGameConfig,
		Clock:    Clock{TimeLeft: DefaultGameConfig.TimePerQuarter},
	}
}
