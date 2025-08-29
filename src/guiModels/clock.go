package guimodels

import (
	"cestoballCounter/src/models"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ClockGuiModel representa la interfaz visual del reloj
type ClockGuiModel struct {
	clock          *models.Clock
	clockRest      *models.Clock
	Container      *fyne.Container
	TimeLabel      *canvas.Text
	TimeLabelRest  *canvas.Text
	ControlButtons *fyne.Container
}

func (c *ClockGuiModel) UpdateLabelTime() {
	c.TimeLabel.Text = formatTime(c.clock.TimeLeft)
	c.TimeLabel.Refresh()
}

func (c *ClockGuiModel) UpdateLabelTimeRest() {
	c.TimeLabelRest.Text = formatTime(c.clockRest.TimeLeft)
	c.TimeLabelRest.Refresh()
}

// GetClock retorna el modelo del reloj
func (c *ClockGuiModel) GetClock() *models.Clock {
	return c.clock
}

// NewClockGuiModel crea una nueva instancia de la interfaz visual del reloj
func NewClockGuiModel(timeLeft time.Duration, timeLeftRest time.Duration) ClockGuiModel {
	clock := &models.Clock{TimeLeft: timeLeft}
	clockRest := &models.Clock{TimeLeft: timeLeftRest}

	clockGuiModel := ClockGuiModel{clock: clock, clockRest: clockRest}
	clock.SetOnUpdate(func() {
		clockGuiModel.UpdateLabelTime()
	})
	clockRest.SetOnUpdate(func() {
		clockGuiModel.UpdateLabelTimeRest()
	})

	clock.SetTimePerQuarter(timeLeft)
	clockRest.SetTimePerQuarter(timeLeftRest)

	clockLabel := NewDefaultText("RELOJ", 40)
	timeLabel := NewDefaultText(formatTime(timeLeft), 80)
	timeLabelRest := NewDefaultText(formatTime(timeLeftRest), 50)
	clockGuiModel.Container = container.NewVBox(clockLabel, layout.NewSpacer(), timeLabel, layout.NewSpacer(), timeLabelRest)
	clockGuiModel.TimeLabel = timeLabel
	clockGuiModel.TimeLabelRest = timeLabelRest

	startButton := widget.NewButton("START", func() {
		clockGuiModel.clock.Start()
	})
	stopButton := widget.NewButton("STOP", func() {
		clockGuiModel.clock.Stop()
	})
	resetButton := widget.NewButton("RESET", func() {
		clockGuiModel.clock.Reset()
	})
	startRestButton := widget.NewButton("MINUTO", func() {
		clockGuiModel.clock.Stop()
		clockGuiModel.clockRest.Start()
	})
	stopRestButton := widget.NewButton("STOP", func() {
		clockGuiModel.clockRest.Reset()
		clockGuiModel.clock.Start()
	})

	controlButtonsLabel := NewDefaultText("RELOJ", 30)
	controlButtons := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer(), stopButton, layout.NewSpacer(), resetButton, layout.NewSpacer())
	controlButtonsRest := container.NewHBox(layout.NewSpacer(), startRestButton, layout.NewSpacer(), stopRestButton, layout.NewSpacer())
	clockGuiModel.ControlButtons = container.NewVBox(controlButtonsLabel, controlButtons, layout.NewSpacer(), controlButtonsRest)

	return clockGuiModel
}

// formatTime formatea el tiempo en formato MM:SS.mm
func formatTime(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	milliseconds := int(d.Milliseconds()) % 1000 / 10 // Obtener centésimas de segundo
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds)
}
