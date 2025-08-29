package viewmodels

import (
	"cestoballCounter/src/observable"
	"time"
)

// TimeOutClockViewModel representa el ViewModel del reloj de TimeOut
type TimeOutClockViewModel struct {
	TimeLeft          *observable.ObservableDuration
	IsRunning         *observable.ObservableBool
	IsStopped         *observable.ObservableBool
	TimeOutTime       time.Duration
	ticker            *time.Ticker
	stopChan          chan bool
	onTimeOutFinished func() // Callback cuando termina el TimeOut
}

// NewTimeOutClockViewModel crea un nuevo ViewModel para el reloj de TimeOut
func NewTimeOutClockViewModel(timeOutTime time.Duration) *TimeOutClockViewModel {
	return &TimeOutClockViewModel{
		TimeLeft:          observable.NewObservableDuration(timeOutTime),
		IsRunning:         observable.NewObservableBool(false),
		IsStopped:         observable.NewObservableBool(false),
		TimeOutTime:       timeOutTime,
		stopChan:          make(chan bool),
		onTimeOutFinished: nil,
	}
}

// SetOnTimeOutFinished establece el callback cuando termina el TimeOut
func (t *TimeOutClockViewModel) SetOnTimeOutFinished(callback func()) {
	t.onTimeOutFinished = callback
}

// Start inicia el reloj de TimeOut
func (t *TimeOutClockViewModel) Start() {
	if t.IsRunning.Get() {
		return
	}

	t.IsRunning.Set(true)
	t.IsStopped.Set(false)

	// Usar un ticker para mayor precisión
	t.ticker = time.NewTicker(50 * time.Millisecond)
	go func() {
		defer t.ticker.Stop()

		for {
			select {
			case <-t.ticker.C:
				if t.IsStopped.Get() {
					break
				}

				currentTime := t.TimeLeft.Get().(time.Duration)
				if currentTime <= 0 {
					t.IsRunning.Set(false)
					// Ejecutar callback cuando termina el TimeOut
					if t.onTimeOutFinished != nil {
						t.onTimeOutFinished()
					}
					return
				}

				delta := 50 * time.Millisecond
				t.TimeLeft.Set(currentTime - delta)

			case <-t.stopChan:
				return
			}
		}
	}()
}

// Stop detiene el reloj de TimeOut
func (t *TimeOutClockViewModel) Stop() {
	t.IsStopped.Set(true)
	t.IsRunning.Set(false)
	if t.ticker != nil {
		t.ticker.Stop()
	}
	select {
	case t.stopChan <- true:
	default:
	}
}

// Reset reinicia el reloj de TimeOut
func (t *TimeOutClockViewModel) Reset() {
	t.TimeLeft.Set(t.TimeOutTime)
	t.IsStopped.Set(true)
	t.IsRunning.Set(false)
	if t.ticker != nil {
		t.ticker.Stop()
	}
}

// GetTimeLeft retorna el tiempo restante
func (t *TimeOutClockViewModel) GetTimeLeft() time.Duration {
	return t.TimeLeft.Get().(time.Duration)
}

// SetTimeOutTime establece el tiempo de TimeOut
func (t *TimeOutClockViewModel) SetTimeOutTime(duration time.Duration) {
	t.TimeOutTime = duration
}

// IsTimeOutFinished retorna si el TimeOut ha terminado
func (t *TimeOutClockViewModel) IsTimeOutFinished() bool {
	return t.TimeLeft.Get().(time.Duration) <= 0
}
