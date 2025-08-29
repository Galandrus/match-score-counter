package viewmodels

import (
	"cestoballCounter/src/observable"
)

// TeamViewModel representa el ViewModel de un equipo
type TeamViewModel struct {
	Name  *observable.Observable
	Score *observable.ObservableInt
}

// NewTeamViewModel crea un nuevo ViewModel para un equipo
func NewTeamViewModel(name string, initialScore int) *TeamViewModel {
	return &TeamViewModel{
		Name:  observable.NewObservable(name),
		Score: observable.NewObservableInt(initialScore),
	}
}

// AddScore añade puntos al equipo
func (t *TeamViewModel) AddScore(score int) {
	t.Score.Set(t.Score.Get() + score)
}

// ResetScore reinicia el puntaje del equipo
func (t *TeamViewModel) ResetScore() {
	t.Score.Set(0)
}

// GetScore retorna el puntaje actual
func (t *TeamViewModel) GetScore() int {
	return t.Score.Get()
}

// GetName retorna el nombre del equipo
func (t *TeamViewModel) GetName() string {
	return t.Name.Get().(string)
}

// SetName establece el nombre del equipo
func (t *TeamViewModel) SetName(name string) {
	t.Name.Set(name)
}
