package observable

import (
	"sync"
)

// Observer es una función que se ejecuta cuando cambia un valor observable
type Observer func(interface{})

// Observable representa un valor que puede ser observado
type Observable struct {
	value     interface{}
	observers []Observer
	mutex     sync.RWMutex
}

// NewObservable crea un nuevo observable con un valor inicial
func NewObservable(initialValue interface{}) *Observable {
	return &Observable{
		value:     initialValue,
		observers: make([]Observer, 0),
	}
}

// Get retorna el valor actual del observable
func (o *Observable) Get() interface{} {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	return o.value
}

// Set establece un nuevo valor y notifica a todos los observadores
func (o *Observable) Set(value interface{}) {
	o.mutex.Lock()
	o.value = value
	observers := make([]Observer, len(o.observers))
	copy(observers, o.observers)
	o.mutex.Unlock()

	// Notificar a todos los observadores
	for _, observer := range observers {
		observer(value)
	}
}

// Subscribe añade un nuevo observador
func (o *Observable) Subscribe(observer Observer) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.observers = append(o.observers, observer)
}

// Unsubscribe remueve un observador específico
func (o *Observable) Unsubscribe(observer Observer) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for i, obs := range o.observers {
		if &obs == &observer {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
			break
		}
	}
}

// ObservableInt es un observable específico para enteros
type ObservableInt struct {
	*Observable
}

// NewObservableInt crea un nuevo observable de enteros
func NewObservableInt(initialValue int) *ObservableInt {
	return &ObservableInt{
		Observable: NewObservable(initialValue),
	}
}

// Get retorna el valor entero actual
func (o *ObservableInt) Get() int {
	return o.Observable.Get().(int)
}

// Set establece un nuevo valor entero
func (o *ObservableInt) Set(value int) {
	o.Observable.Set(value)
}

// ObservableDuration es un observable específico para duraciones
type ObservableDuration struct {
	*Observable
}

// NewObservableDuration crea un nuevo observable de duraciones
func NewObservableDuration(initialValue interface{}) *ObservableDuration {
	return &ObservableDuration{
		Observable: NewObservable(initialValue),
	}
}

// Get retorna la duración actual
func (o *ObservableDuration) Get() interface{} {
	return o.Observable.Get()
}

// Set establece una nueva duración
func (o *ObservableDuration) Set(value interface{}) {
	o.Observable.Set(value)
}

// ObservableBool es un observable específico para booleanos
type ObservableBool struct {
	*Observable
}

// NewObservableBool crea un nuevo observable de booleanos
func NewObservableBool(initialValue bool) *ObservableBool {
	return &ObservableBool{
		Observable: NewObservable(initialValue),
	}
}

// Get retorna el valor booleano actual
func (o *ObservableBool) Get() bool {
	return o.Observable.Get().(bool)
}

// Set establece un nuevo valor booleano
func (o *ObservableBool) Set(value bool) {
	o.Observable.Set(value)
}
