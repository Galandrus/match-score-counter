package viewmodels

import (
	"cestoballCounter/src/observable"
	"time"
)

// ClockViewModel representa el ViewModel del reloj
type ClockViewModel struct {
	TimeLeft       *observable.ObservableDuration
	IsRunning      *observable.ObservableBool
	IsStopped      *observable.ObservableBool
	Quarter        *observable.ObservableInt
	TimePerQuarter time.Duration
	ticker         *time.Ticker
	stopChan       chan bool
}

// NewClockViewModel crea un nuevo ViewModel para el reloj
func NewClockViewModel(timePerQuarter time.Duration) *ClockViewModel {
	return &ClockViewModel{
		TimeLeft:       observable.NewObservableDuration(timePerQuarter),
		IsRunning:      observable.NewObservableBool(false),
		IsStopped:      observable.NewObservableBool(false),
		Quarter:        observable.NewObservableInt(1),
		TimePerQuarter: timePerQuarter,
		stopChan:       make(chan bool),
	}
}

// Start inicia el reloj
func (c *ClockViewModel) Start() {
	if c.IsRunning.Get() {
		return
	}

	c.IsRunning.Set(true)
	c.IsStopped.Set(false)

	// Usar un ticker para mayor precisión
	c.ticker = time.NewTicker(50 * time.Millisecond)
	go func() {
		defer c.ticker.Stop()

		for {
			select {
			case <-c.ticker.C:
				if c.IsStopped.Get() {
					break
				}

				currentTime := c.TimeLeft.Get().(time.Duration)
				if currentTime <= 0 {
					c.IsRunning.Set(false)
					return
				}

				delta := 50 * time.Millisecond
				c.TimeLeft.Set(currentTime - delta)

			case <-c.stopChan:
				return
			}
		}
	}()
}

// Stop detiene el reloj
func (c *ClockViewModel) Stop() {
	c.IsStopped.Set(true)
	c.IsRunning.Set(false)
	if c.ticker != nil {
		c.ticker.Stop()
	}
	select {
	case c.stopChan <- true:
	default:
	}
}

// Reset reinicia el reloj
func (c *ClockViewModel) Reset() {
	c.TimeLeft.Set(c.TimePerQuarter)
	c.IsStopped.Set(true)
	c.IsRunning.Set(false)
	if c.ticker != nil {
		c.ticker.Stop()
	}
}

// GetTimeLeft retorna el tiempo restante
func (c *ClockViewModel) GetTimeLeft() time.Duration {
	return c.TimeLeft.Get().(time.Duration)
}

// SetTimePerQuarter establece el tiempo por cuarto
func (c *ClockViewModel) SetTimePerQuarter(duration time.Duration) {
	c.TimePerQuarter = duration
}
