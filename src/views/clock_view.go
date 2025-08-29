package views

import (
	"cestoballCounter/src/viewmodels"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ClockView representa la vista del reloj con data binding automático
type ClockView struct {
	ViewModel      *viewmodels.ClockViewModel
	GameViewModel  *viewmodels.GameViewModel
	Container      *fyne.Container
	TimeLabel      *canvas.Text
	TimeLabelRest  *canvas.Text
	ControlButtons *fyne.Container
}

// NewClockView crea una nueva vista del reloj con data binding automático
func NewClockView(viewModel *viewmodels.ClockViewModel, gameViewModel *viewmodels.GameViewModel) *ClockView {
	clockView := &ClockView{
		ViewModel:     viewModel,
		GameViewModel: gameViewModel,
	}

	// Crear elementos de la UI
	clockLabel := canvas.NewText("RELOJ", nil)
	clockLabel.TextSize = 40
	clockLabel.Alignment = fyne.TextAlignCenter

	timeLabel := canvas.NewText(formatTime(viewModel.GetTimeLeft()), nil)
	timeLabel.TextSize = 80
	timeLabel.Alignment = fyne.TextAlignCenter

	timeLabelRest := canvas.NewText("00:00.00", nil)
	timeLabelRest.TextSize = 50
	timeLabelRest.Alignment = fyne.TextAlignCenter

	clockView.TimeLabel = timeLabel
	clockView.TimeLabelRest = timeLabelRest

	// Crear contenedor principal
	clockContainer := container.NewVBox(clockLabel, layout.NewSpacer(), timeLabel, layout.NewSpacer(), timeLabelRest)
	clockView.Container = clockContainer

	// Crear botones de control
	startButton := widget.NewButton("START", func() {
		viewModel.Start()
	})
	stopButton := widget.NewButton("STOP", func() {
		viewModel.Stop()
	})
	resetButton := widget.NewButton("RESET", func() {
		viewModel.Reset()
	})

	controlButtonsLabel := canvas.NewText("RELOJ", nil)
	controlButtonsLabel.TextSize = 30
	controlButtonsLabel.Alignment = fyne.TextAlignCenter

	controlButtons := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer(), stopButton, layout.NewSpacer(), resetButton, layout.NewSpacer())
	controlButtonsContainer := container.NewVBox(controlButtonsLabel, controlButtons)
	clockView.ControlButtons = controlButtonsContainer

	// Configurar data binding automático
	clockView.setupDataBinding()

	return clockView
}

// setupDataBinding configura el data binding automático
func (c *ClockView) setupDataBinding() {
	// Binding para el tiempo restante
	c.ViewModel.TimeLeft.Subscribe(func(value interface{}) {
		duration := value.(time.Duration)
		c.TimeLabel.Text = formatTime(duration)
		c.TimeLabel.Refresh()
	})

	// Binding para el estado del reloj
	c.ViewModel.IsRunning.Subscribe(func(value interface{}) {
		// TODO: Actualizar estado visual de los botones si es necesario
	})

	c.ViewModel.IsStopped.Subscribe(func(value interface{}) {
		// TODO: Actualizar estado visual de los botones si es necesario
	})
}

// GetContainer retorna el contenedor del reloj
func (c *ClockView) GetContainer() *fyne.Container {
	return c.Container
}

// GetControlButtons retorna los botones de control
func (c *ClockView) GetControlButtons() *fyne.Container {
	return c.ControlButtons
}

// formatTime formatea el tiempo en formato MM:SS.mm
func formatTime(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	milliseconds := int(d.Milliseconds()) % 1000 / 10 // Obtener centésimas de segundo
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds)
}
