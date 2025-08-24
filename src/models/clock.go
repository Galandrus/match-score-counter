package models

import "time"

// Clock representa la lógica del reloj del juego
type Clock struct {
	timePerQuarter time.Duration
	Quarter        int
	TimeLeft       time.Duration
	IsRunning      bool
	IsStopped      bool
	onUpdate       func()
}

// Start inicia el reloj
func (c *Clock) Start() {
	if c.IsRunning {
		return
	}

	c.IsRunning = true
	c.IsStopped = false
	for c.TimeLeft > 0 {
		delta := 10 * time.Millisecond
		c.UpdateTime(delta)
		if c.onUpdate != nil {
			c.onUpdate()
		}
		time.Sleep(delta)
		if c.IsStopped {
			c.IsRunning = false
			break
		}
	}
}

// Stop detiene el reloj
func (c *Clock) Stop() {
	c.IsStopped = true
	c.IsRunning = false
}

// Reset reinicia el reloj
func (c *Clock) Reset() {
	c.TimeLeft = c.timePerQuarter
	c.IsStopped = false
	c.IsRunning = false

	if c.onUpdate != nil {
		c.onUpdate()
	}
}

func (c *Clock) SetOnUpdate(onUpdate func()) {
	c.onUpdate = onUpdate
}

func (c *Clock) SetTimePerQuarter(timePerQuarter time.Duration) {
	c.timePerQuarter = timePerQuarter
}

// UpdateTime actualiza el tiempo restante
func (c *Clock) UpdateTime(delta time.Duration) {
	c.TimeLeft -= delta
	if c.TimeLeft < 0 {
		c.TimeLeft = 0
	}
}
