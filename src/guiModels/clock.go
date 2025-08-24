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
	Container      *fyne.Container
	TimeLabel      *canvas.Text
	ControlButtons *fyne.Container
}

func (c *ClockGuiModel) UpdateLabelTime() {
	c.TimeLabel.Text = formatTime(c.clock.TimeLeft)
	c.TimeLabel.Refresh()
}

// GetClock retorna el modelo del reloj
func (c *ClockGuiModel) GetClock() *models.Clock {
	return c.clock
}

// NewClockGuiModel crea una nueva instancia de la interfaz visual del reloj
func NewClockGuiModel(timeLeft time.Duration) ClockGuiModel {
	clock := &models.Clock{TimeLeft: timeLeft}
	clockGuiModel := ClockGuiModel{clock: clock}
	clock.SetOnUpdate(func() {
		clockGuiModel.UpdateLabelTime()
	})
	clock.SetTimePerQuarter(timeLeft)

	clockLabel := NewDefaultText("RELOJ", 40)
	timeLabel := NewDefaultText(formatTime(timeLeft), 80)
	clockGuiModel.Container = container.NewVBox(clockLabel, layout.NewSpacer(), timeLabel)
	clockGuiModel.TimeLabel = timeLabel

	startButton := widget.NewButton("START", func() {
		go clockGuiModel.clock.Start()
	})
	stopButton := widget.NewButton("STOP", func() {
		go clockGuiModel.clock.Stop()
	})
	resetButton := widget.NewButton("RESET", func() {
		go clockGuiModel.clock.Reset()
	})

	controlButtonsLabel := NewDefaultText("RELOJ", 30)
	controlButtons := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer(), stopButton, layout.NewSpacer(), resetButton, layout.NewSpacer())
	clockGuiModel.ControlButtons = container.NewVBox(controlButtonsLabel, controlButtons, layout.NewSpacer())

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
