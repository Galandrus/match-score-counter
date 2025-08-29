package viewmodels

import (
	"cestoballCounter/src/observable"
	"time"
)

// RestClockViewModel representa el ViewModel del reloj de descanso
type RestClockViewModel struct {
	TimeLeft       *observable.ObservableDuration
	IsRunning      *observable.ObservableBool
	IsStopped      *observable.ObservableBool
	RestTime       time.Duration
	ticker         *time.Ticker
	stopChan       chan bool
	onRestFinished func() // Callback cuando termina el descanso
}

// NewRestClockViewModel crea un nuevo ViewModel para el reloj de descanso
func NewRestClockViewModel(restTime time.Duration) *RestClockViewModel {
	return &RestClockViewModel{
		TimeLeft:       observable.NewObservableDuration(restTime),
		IsRunning:      observable.NewObservableBool(false),
		IsStopped:      observable.NewObservableBool(false),
		RestTime:       restTime,
		stopChan:       make(chan bool),
		onRestFinished: nil,
	}
}

// SetOnRestFinished establece el callback cuando termina el descanso
func (r *RestClockViewModel) SetOnRestFinished(callback func()) {
	r.onRestFinished = callback
}

// Start inicia el reloj de descanso
func (r *RestClockViewModel) Start() {
	if r.IsRunning.Get() {
		return
	}

	r.IsRunning.Set(true)
	r.IsStopped.Set(false)

	// Usar un ticker para mayor precisión
	r.ticker = time.NewTicker(20 * time.Millisecond)
	go func() {
		defer r.ticker.Stop()

		for {
			select {
			case <-r.ticker.C:
				if r.IsStopped.Get() {
					break
				}

				currentTime := r.TimeLeft.Get().(time.Duration)
				if currentTime <= 0 {
					r.IsRunning.Set(false)
					// Ejecutar callback cuando termina el descanso
					if r.onRestFinished != nil {
						r.onRestFinished()
					}
					return
				}

				delta := 50 * time.Millisecond
				r.TimeLeft.Set(currentTime - delta)

			case <-r.stopChan:
				return
			}
		}
	}()
}

// Stop detiene el reloj de descanso
func (r *RestClockViewModel) Stop() {
	r.IsStopped.Set(true)
	r.IsRunning.Set(false)
	if r.ticker != nil {
		r.ticker.Stop()
	}
	select {
	case r.stopChan <- true:
	default:
	}
}

// Reset reinicia el reloj de descanso
func (r *RestClockViewModel) Reset() {
	r.TimeLeft.Set(r.RestTime)
	r.IsStopped.Set(true)
	r.IsRunning.Set(false)
	if r.ticker != nil {
		r.ticker.Stop()
	}
}

// GetTimeLeft retorna el tiempo restante
func (r *RestClockViewModel) GetTimeLeft() time.Duration {
	return r.TimeLeft.Get().(time.Duration)
}

// SetRestTime establece el tiempo de descanso
func (r *RestClockViewModel) SetRestTime(duration time.Duration) {
	r.RestTime = duration
}

// IsRestFinished retorna si el descanso ha terminado
func (r *RestClockViewModel) IsRestFinished() bool {
	return r.TimeLeft.Get().(time.Duration) <= 0
}
