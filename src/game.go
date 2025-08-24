package src

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

type Team struct {
	Name  string
	Score int
}

func (t *Team) AddScore(score int) {
	t.Score += score
}

type Game struct {
	HomeTeam TeamGuiModel
	AwayTeam TeamGuiModel
	Config   GameConfig
	TimeLeft time.Duration
	Clock    ClockGuiModel
}

func (g *Game) Start() {
	g.Clock.Start()
}

func NewGame() *Game {
	return &Game{
		HomeTeam: NewTeamGuiModel(Team{Name: "Local", Score: 0}),
		AwayTeam: NewTeamGuiModel(Team{Name: "Visita", Score: 0}),
		Config:   DefaultGameConfig,
		TimeLeft: DefaultGameConfig.TimePerQuarter,
		Clock:    NewClockGuiModel(DefaultGameConfig.TimePerQuarter),
	}
}
