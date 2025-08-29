package viewmodels

import (
	"time"
)

// GameConfig contiene la configuración del juego
type GameConfig struct {
	Quarters         int
	TimePerQuarter   time.Duration
	RestTime         time.Duration
	TimePerTimeOut   time.Duration
	TimePerExtraTime time.Duration
}

// DefaultGameConfig configuración por defecto del juego
var DefaultGameConfig = GameConfig{
	Quarters:         2,
	TimePerQuarter:   20 * time.Minute,
	RestTime:         5 * time.Minute,
	TimePerTimeOut:   1 * time.Minute,
	TimePerExtraTime: 3 * time.Minute,
}

// GameViewModel representa el ViewModel principal del juego
type GameViewModel struct {
	HomeTeam     *TeamViewModel
	AwayTeam     *TeamViewModel
	Clock        *ClockViewModel
	RestClock    *RestClockViewModel
	TimeOutClock *TimeOutClockViewModel
	Config       GameConfig
}

// NewGameViewModel crea un nuevo ViewModel del juego
func NewGameViewModel() *GameViewModel {
	gameVM := &GameViewModel{
		HomeTeam:     NewTeamViewModel("Local", 0),
		AwayTeam:     NewTeamViewModel("Visita", 0),
		Clock:        NewClockViewModel(DefaultGameConfig.TimePerQuarter),
		RestClock:    NewRestClockViewModel(DefaultGameConfig.RestTime),
		TimeOutClock: NewTimeOutClockViewModel(DefaultGameConfig.TimePerTimeOut),
		Config:       DefaultGameConfig,
	}

	// Configurar callback para cuando termina el descanso
	gameVM.RestClock.SetOnRestFinished(func() {
		// Cuando termina el descanso, reiniciar el reloj principal
		gameVM.Clock.Reset()
		gameVM.Clock.Start()
	})

	// Configurar callback para cuando termina el TimeOut
	gameVM.TimeOutClock.SetOnTimeOutFinished(func() {
		// Cuando termina el TimeOut, reiniciar el reloj principal
		gameVM.Clock.Reset()
		gameVM.Clock.Start()
	})

	// Configurar callback para cuando termina un tiempo
	gameVM.Clock.SetOnQuarterFinished(func() {
		// Cuando termina un tiempo, avanzar al siguiente
		gameVM.Clock.NextQuarter()
	})

	return gameVM
}

// NewGameViewModelWithConfig crea un nuevo ViewModel del juego con configuración personalizada
func NewGameViewModelWithConfig(config GameConfig) *GameViewModel {
	gameVM := &GameViewModel{
		HomeTeam:     NewTeamViewModel("Local", 0),
		AwayTeam:     NewTeamViewModel("Visita", 0),
		Clock:        NewClockViewModel(config.TimePerQuarter),
		RestClock:    NewRestClockViewModel(config.RestTime),
		TimeOutClock: NewTimeOutClockViewModel(config.TimePerTimeOut),
		Config:       config,
	}

	// Configurar callback para cuando termina el descanso
	gameVM.RestClock.SetOnRestFinished(func() {
		gameVM.RestClock.Reset()
	})

	// Configurar callback para cuando termina el TimeOut
	gameVM.TimeOutClock.SetOnTimeOutFinished(func() {
		gameVM.TimeOutClock.Reset()
	})

	// Configurar callback para cuando termina un tiempo
	gameVM.Clock.SetOnQuarterFinished(func() {
		// Cuando termina un tiempo, avanzar al siguiente
		gameVM.Clock.NextQuarter()
	})

	return gameVM
}

// StartGame inicia el juego
func (g *GameViewModel) StartGame() {
	g.Clock.Start()
}

// StopGame detiene el juego
func (g *GameViewModel) StopGame() {
	g.Clock.Stop()
}

// ResetGame reinicia el juego
func (g *GameViewModel) ResetGame() {
	g.HomeTeam.ResetScore()
	g.AwayTeam.ResetScore()
	g.Clock.Reset()
}

// AddScoreToHomeTeam añade puntos al equipo local
func (g *GameViewModel) AddScoreToHomeTeam(score int) {
	g.HomeTeam.AddScore(score)
}

// AddScoreToAwayTeam añade puntos al equipo visitante
func (g *GameViewModel) AddScoreToAwayTeam(score int) {
	g.AwayTeam.AddScore(score)
}

// GetHomeTeamScore retorna el puntaje del equipo local
func (g *GameViewModel) GetHomeTeamScore() int {
	return g.HomeTeam.GetScore()
}

// GetAwayTeamScore retorna el puntaje del equipo visitante
func (g *GameViewModel) GetAwayTeamScore() int {
	return g.AwayTeam.GetScore()
}

// GetTimeLeft retorna el tiempo restante
func (g *GameViewModel) GetTimeLeft() time.Duration {
	return g.Clock.GetTimeLeft()
}

// IsClockRunning retorna si el reloj está corriendo
func (g *GameViewModel) IsClockRunning() bool {
	return g.Clock.IsRunning.Get()
}

// IsClockStopped retorna si el reloj está detenido
func (g *GameViewModel) IsClockStopped() bool {
	return g.Clock.IsStopped.Get()
}

// StartRest inicia el reloj de descanso
func (g *GameViewModel) StartRest() {
	g.Clock.Stop()
	g.RestClock.Start()
}

// StopRest detiene el reloj de descanso
func (g *GameViewModel) StopRest() {
	g.RestClock.Stop()
	g.Clock.Start()
}

// GetRestTimeLeft retorna el tiempo restante del descanso
func (g *GameViewModel) GetRestTimeLeft() time.Duration {
	return g.RestClock.GetTimeLeft()
}

// IsRestRunning retorna si el reloj de descanso está corriendo
func (g *GameViewModel) IsRestRunning() bool {
	return g.RestClock.IsRunning.Get()
}

// IsRestStopped retorna si el reloj de descanso está detenido
func (g *GameViewModel) IsRestStopped() bool {
	return g.RestClock.IsStopped.Get()
}

// StartTimeOut inicia el reloj de TimeOut
func (g *GameViewModel) StartTimeOut() {
	g.Clock.Stop()
	g.TimeOutClock.Start()
}

// StopTimeOut detiene el reloj de TimeOut
func (g *GameViewModel) StopTimeOut() {
	g.TimeOutClock.Stop()
	g.Clock.Start()
}

// GetTimeOutTimeLeft retorna el tiempo restante del TimeOut
func (g *GameViewModel) GetTimeOutTimeLeft() time.Duration {
	return g.TimeOutClock.GetTimeLeft()
}

// IsTimeOutRunning retorna si el reloj de TimeOut está corriendo
func (g *GameViewModel) IsTimeOutRunning() bool {
	return g.TimeOutClock.IsRunning.Get()
}

// IsTimeOutStopped retorna si el reloj de TimeOut está detenido
func (g *GameViewModel) IsTimeOutStopped() bool {
	return g.TimeOutClock.IsStopped.Get()
}

// GetCurrentQuarter retorna el tiempo actual
func (g *GameViewModel) GetCurrentQuarter() int {
	return g.Clock.GetQuarter()
}

// NextQuarter avanza al siguiente tiempo
func (g *GameViewModel) NextQuarter() {
	g.Clock.NextQuarter()
}

// IsGameFinished retorna si el juego ha terminado
func (g *GameViewModel) IsGameFinished() bool {
	return g.Clock.IsGameFinished()
}
