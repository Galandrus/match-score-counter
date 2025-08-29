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

// TimeOutClockView representa la vista del reloj de TimeOut con data binding automático
type TimeOutClockView struct {
	ViewModel      *viewmodels.TimeOutClockViewModel
	GameViewModel  *viewmodels.GameViewModel
	Container      *fyne.Container
	TimeLabel      *canvas.Text
	ControlButtons *fyne.Container
}

// NewTimeOutClockView crea una nueva vista del reloj de TimeOut con data binding automático
func NewTimeOutClockView(viewModel *viewmodels.TimeOutClockViewModel, gameViewModel *viewmodels.GameViewModel) *TimeOutClockView {
	timeOutClockView := &TimeOutClockView{
		ViewModel:     viewModel,
		GameViewModel: gameViewModel,
	}

	// Crear elementos de la UI
	timeOutLabel := guimodels.NewDefaultText("Minuto", 30)

	timeLabel := guimodels.NewDefaultText(formatTimeOutTime(viewModel.GetTimeLeft()), 40)

	timeOutClockView.TimeLabel = timeLabel

	// Crear contenedor principal
	timeOutContainer := container.NewHBox(timeOutLabel, layout.NewSpacer(), timeLabel, layout.NewSpacer())
	timeOutClockView.Container = timeOutContainer

	// Crear botones de control
	startTimeOutButton := widget.NewButton("Iniciar Minuto", func() {
		gameViewModel.StartTimeOut()
	})
	stopTimeOutButton := widget.NewButton("Detener Minuto", func() {
		gameViewModel.StopTimeOut()
		viewModel.Reset()
	})
	resetTimeOutButton := widget.NewButton("Reiniciar Minuto", func() {
		viewModel.Reset()
	})

	controlButtons := container.NewHBox(
		layout.NewSpacer(),
		startTimeOutButton,
		layout.NewSpacer(),
		stopTimeOutButton,
		layout.NewSpacer(),
		resetTimeOutButton,
		layout.NewSpacer(),
	)
	controlButtonsLabel := guimodels.NewDefaultText("MINUTO", 30)
	timeOutClockView.ControlButtons = container.NewVBox(controlButtonsLabel, controlButtons, layout.NewSpacer())

	// Configurar data binding automático
	timeOutClockView.setupDataBinding()

	return timeOutClockView
}

// setupDataBinding configura el data binding automático
func (t *TimeOutClockView) setupDataBinding() {
	// Binding para el tiempo restante
	t.ViewModel.TimeLeft.Subscribe(func(value interface{}) {
		duration := value.(time.Duration)
		t.TimeLabel.Text = formatTimeOutTime(duration)
		t.TimeLabel.Refresh()
	})

	// Binding para el estado del reloj
	t.ViewModel.IsRunning.Subscribe(func(value interface{}) {
		// TODO: Actualizar estado visual de los botones si es necesario
	})

	t.ViewModel.IsStopped.Subscribe(func(value interface{}) {
		// TODO: Actualizar estado visual de los botones si es necesario
	})
}

// GetContainer retorna el contenedor del reloj de TimeOut
func (t *TimeOutClockView) GetContainer() *fyne.Container {
	return t.Container
}

// GetControlButtons retorna los botones de control
func (t *TimeOutClockView) GetControlButtons() *fyne.Container {
	return t.ControlButtons
}

// GetDisplayContainer retorna solo el contenedor de visualización (sin controles)
func (t *TimeOutClockView) GetDisplayContainer() *fyne.Container {
	return t.Container
}

// formatTimeOutTime formatea el tiempo en formato MM:SS.mm
func formatTimeOutTime(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	milliseconds := int(d.Milliseconds()) % 1000 / 10 // Obtener centésimas de segundo
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds)
}
