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

// RestClockView representa la vista del reloj de descanso con data binding automático
type RestClockView struct {
	ViewModel      *viewmodels.RestClockViewModel
	Container      *fyne.Container
	TimeLabel      *canvas.Text
	ControlButtons *fyne.Container
}

// NewRestClockView crea una nueva vista del reloj de descanso con data binding automático
func NewRestClockView(viewModel *viewmodels.RestClockViewModel) *RestClockView {
	restClockView := &RestClockView{
		ViewModel: viewModel,
	}

	// Crear elementos de la UI
	restLabel := canvas.NewText("DESCANSO", nil)
	restLabel.TextSize = 30
	restLabel.Alignment = fyne.TextAlignCenter

	timeLabel := canvas.NewText(formatRestTime(viewModel.GetTimeLeft()), nil)
	timeLabel.TextSize = 60
	timeLabel.Alignment = fyne.TextAlignCenter

	restClockView.TimeLabel = timeLabel

	// Crear contenedor principal
	restContainer := container.NewVBox(restLabel, layout.NewSpacer(), timeLabel, layout.NewSpacer())
	restClockView.Container = restContainer

	// Crear botones de control
	startRestButton := widget.NewButton("INICIAR DESCANSO", func() {
		viewModel.Start()
	})
	stopRestButton := widget.NewButton("TERMINAR DESCANSO", func() {
		viewModel.Stop()
	})
	resetRestButton := widget.NewButton("RESET", func() {
		viewModel.Reset()
	})

	controlButtons := container.NewHBox(
		layout.NewSpacer(),
		startRestButton,
		layout.NewSpacer(),
		stopRestButton,
		layout.NewSpacer(),
		resetRestButton,
		layout.NewSpacer(),
	)
	restClockView.ControlButtons = controlButtons

	// Configurar data binding automático
	restClockView.setupDataBinding()

	return restClockView
}

// setupDataBinding configura el data binding automático
func (r *RestClockView) setupDataBinding() {
	// Binding para el tiempo restante
	r.ViewModel.TimeLeft.Subscribe(func(value interface{}) {
		duration := value.(time.Duration)
		r.TimeLabel.Text = formatRestTime(duration)
		r.TimeLabel.Refresh()
	})

	// Binding para el estado del reloj
	r.ViewModel.IsRunning.Subscribe(func(value interface{}) {
		// TODO: Actualizar estado visual de los botones si es necesario
	})

	r.ViewModel.IsStopped.Subscribe(func(value interface{}) {
		// TODO: Actualizar estado visual de los botones si es necesario
	})
}

// GetContainer retorna el contenedor del reloj de descanso
func (r *RestClockView) GetContainer() *fyne.Container {
	return r.Container
}

// GetControlButtons retorna los botones de control
func (r *RestClockView) GetControlButtons() *fyne.Container {
	return r.ControlButtons
}

// formatRestTime formatea el tiempo en formato MM:SS.mm
func formatRestTime(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	milliseconds := int(d.Milliseconds()) % 1000 / 10 // Obtener centésimas de segundo
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds)
}
