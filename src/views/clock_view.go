package views

import (
	guimodels "cestoballCounter/src/guiModels"
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
	QuarterLabel   *canvas.Text
	ControlButtons *fyne.Container
}

// NewClockView crea una nueva vista del reloj con data binding automático
func NewClockView(viewModel *viewmodels.ClockViewModel, gameViewModel *viewmodels.GameViewModel) *ClockView {
	clockView := &ClockView{
		ViewModel:     viewModel,
		GameViewModel: gameViewModel,
	}

	// Crear elementos de la UI
	quarterLabel := guimodels.NewDefaultText(getQuarterText(viewModel.GetQuarter()), 40)
	timeLabel := guimodels.NewDefaultText(formatTime(viewModel.GetTimeLeft()), 80)

	clockView.TimeLabel = timeLabel
	clockView.QuarterLabel = quarterLabel

	// Crear contenedor principal
	clockContainer := container.NewVBox(quarterLabel, layout.NewSpacer(), timeLabel, layout.NewSpacer())
	clockView.Container = clockContainer

	// Crear botones de control
	startButton := widget.NewButton("Iniciar Tiempo", func() {
		viewModel.Start()
	})
	stopButton := widget.NewButton("Detener Tiempo", func() {
		viewModel.Stop()
	})
	resetButton := widget.NewButton("Reiniciar Tiempo", func() {
		viewModel.Reset()
	})

	controlButtonsLabel := guimodels.NewDefaultText("TIEMPO", 30)

	controlButtons := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer(), stopButton, layout.NewSpacer(), resetButton, layout.NewSpacer())
	controlButtonsContainer := container.NewVBox(controlButtonsLabel, controlButtons, layout.NewSpacer())
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

	// Binding para el tiempo actual
	c.ViewModel.Quarter.Subscribe(func(value interface{}) {
		quarter := value.(int)
		c.QuarterLabel.Text = getQuarterText(quarter)
		c.QuarterLabel.Refresh()
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

// GetDisplayContainer retorna solo el contenedor de visualización (sin controles)
func (c *ClockView) GetDisplayContainer() *fyne.Container {
	return c.Container
}

// formatTime formatea el tiempo en formato MM:SS.mm
func formatTime(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	milliseconds := int(d.Milliseconds()) % 1000 / 10 // Obtener centésimas de segundo
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds)
}

// getQuarterText retorna el texto del tiempo actual
func getQuarterText(quarter int) string {
	switch quarter {
	case 1:
		return "PRIMER TIEMPO"
	case 2:
		return "SEGUNDO TIEMPO"
	default:
		return "FIN DEL JUEGO"
	}
}
