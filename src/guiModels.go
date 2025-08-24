package src

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type TeamGuiModel struct {
	Team
	TeamLabel     *canvas.Text
	ScoreLabel    *canvas.Text
	TeamContainer *fyne.Container
	ScoreButtons  *fyne.Container
}

func (t *TeamGuiModel) AddScore(score int) {
	t.Team.AddScore(score)
	t.ScoreLabel.Text = strconv.Itoa(t.Team.Score)
	t.ScoreLabel.Refresh()
}

func NewTeamGuiModel(team Team) TeamGuiModel {
	teamGuiModel := TeamGuiModel{Team: team}
	teamLabel := NewDefaultText(team.Name, 40)
	scoreLabel := NewDefaultText(strconv.Itoa(team.Score), 40)
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

type GameGuiModel struct {
	HomeTeam   TeamGuiModel
	AwayTeam   TeamGuiModel
	ClockLabel *widget.Label
}

type ClockGuiModel struct {
	Quarter        int
	TimeLeft       time.Duration
	Container      *fyne.Container
	TimeLabel      *canvas.Text
	isRunning      bool
	stop           bool
	ControlButtons *fyne.Container
}

func (c *ClockGuiModel) Start() {
	if c.isRunning {
		return
	}

	c.isRunning = true
	for c.TimeLeft > 0 {
		c.TimeLeft -= 10 * time.Millisecond
		go c.UpdateTime()
		time.Sleep(10 * time.Millisecond)
		if c.stop {
			c.isRunning = false
			break
		}
	}
}

func (c *ClockGuiModel) Stop() {
	c.stop = true
}

func (c *ClockGuiModel) Reset() {
	c.TimeLeft = DefaultGameConfig.TimePerQuarter
	c.UpdateTime()
	c.Stop()
}

func (c *ClockGuiModel) Resume() {
	c.stop = false
	c.Start()
}

func (c *ClockGuiModel) UpdateTime() {
	c.TimeLabel.Text = formatTime(c.TimeLeft)
	c.TimeLabel.Refresh()
}

func NewClockGuiModel(timeLeft time.Duration) ClockGuiModel {
	clockGuiModel := ClockGuiModel{TimeLeft: timeLeft}
	clockLabel := NewDefaultText("RELOJ", 40)
	timeLabel := NewDefaultText(formatTime(timeLeft), 80)
	clockGuiModel.Container = container.NewVBox(clockLabel, layout.NewSpacer(), timeLabel)
	clockGuiModel.TimeLabel = timeLabel

	startButton := widget.NewButton("START", func() {
		go clockGuiModel.Resume()
	})
	stopButton := widget.NewButton("STOP", func() {
		go clockGuiModel.Stop()
	})
	resetButton := widget.NewButton("RESET", func() {
		go clockGuiModel.Reset()
	})

	controlButtonsLabel := NewDefaultText("RELOJ", 30)
	controlButtons := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer(), stopButton, layout.NewSpacer(), resetButton, layout.NewSpacer())
	clockGuiModel.ControlButtons = container.NewVBox(controlButtonsLabel, controlButtons, layout.NewSpacer())

	return clockGuiModel
}

func NewDefaultText(text string, size float32) *canvas.Text {
	textCanvas := canvas.NewText(text, nil)
	textCanvas.TextSize = size
	textCanvas.TextStyle = fyne.TextStyle{Bold: true}
	textCanvas.Alignment = fyne.TextAlignCenter

	return textCanvas
}
