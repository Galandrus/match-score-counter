package models

// Team representa un equipo con su lógica de negocio pura
type Team struct {
	Name  string
	Score int
}

// AddScore añade puntos al equipo
func (t *Team) AddScore(score int) {
	t.Score += score
}

// ResetScore reinicia el puntaje del equipo
func (t *Team) ResetScore() {
	t.Score = 0
}

// GetScore retorna el puntaje actual
func (t *Team) GetScore() int {
	return t.Score
}

// GetName retorna el nombre del equipo
func (t *Team) GetName() string {
	return t.Name
}
