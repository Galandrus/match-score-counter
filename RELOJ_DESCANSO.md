# 🕐 Reloj de Descanso - Funcionalidad Implementada

## 🎯 Descripción

Se ha implementado un **reloj de descanso separado** que funciona de manera independiente al reloj principal del juego. Este reloj cuenta un minuto y automáticamente reinicia el reloj principal cuando termina.

## 🏗️ Arquitectura Implementada

### 1. **RestClockViewModel** (`src/viewmodels/rest_clock_viewmodel.go`)

```go
type RestClockViewModel struct {
    TimeLeft       *observable.ObservableDuration
    IsRunning      *observable.ObservableBool
    IsStopped      *observable.ObservableBool
    RestTime       time.Duration
    ticker         *time.Ticker
    stopChan       chan bool
    onRestFinished func() // Callback cuando termina el descanso
}
```

**Características:**

- ✅ **Reloj independiente** con su propio ticker
- ✅ **Data binding automático** con observables
- ✅ **Callback automático** cuando termina el descanso
- ✅ **Thread-safe** con canales de comunicación

### 2. **RestClockView** (`src/views/rest_clock_view.go`)

```go
type RestClockView struct {
    ViewModel      *viewmodels.RestClockViewModel
    Container      *fyne.Container
    TimeLabel      *canvas.Text
    ControlButtons *fyne.Container
}
```

**Características:**

- ✅ **Vista separada** del reloj de descanso
- ✅ **Data binding automático** con el ViewModel
- ✅ **Botones de control** específicos
- ✅ **Actualización automática** del display

## 🔄 Flujo de Funcionamiento

### **Secuencia Normal:**

1. **Reloj principal** está corriendo
2. **Usuario** hace clic en "INICIAR DESCANSO"
3. **Reloj principal** se detiene automáticamente
4. **Reloj de descanso** inicia la cuenta regresiva (1 minuto)
5. **Cuando termina** el descanso:
   - Se ejecuta el callback automáticamente
   - **Reloj principal** se reinicia
   - **Reloj principal** inicia automáticamente
6. **Reloj de descanso** se resetea

### **Control Manual:**

- **Usuario** puede hacer clic en "TERMINAR DESCANSO"
- **Reloj de descanso** se detiene
- **Reloj principal** inicia inmediatamente

## 🎮 Controles Disponibles

### **Reloj Principal:**

- **START**: Inicia el reloj principal
- **STOP**: Detiene el reloj principal
- **RESET**: Reinicia el reloj principal
- **INICIAR DESCANSO**: Detiene el principal e inicia el descanso

### **Reloj de Descanso:**

- **INICIAR DESCANSO**: Inicia el reloj de descanso
- **TERMINAR DESCANSO**: Detiene el descanso e inicia el principal
- **RESET**: Reinicia el reloj de descanso

## 📊 Integración con GameViewModel

### **Nuevos Métodos Agregados:**

```go
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
```

### **Configuración Automática:**

```go
// En NewGameViewModel()
gameVM.RestClock.SetOnRestFinished(func() {
    // Cuando termina el descanso, reiniciar el reloj principal
    gameVM.Clock.Reset()
    gameVM.Clock.Start()
})
```

## 🎨 Interfaz de Usuario

### **Layout Actualizado:**

```
┌─────────────────────────────────────┐
│              MARCADOR               │
├─────────────────────────────────────┤
│    LOCAL     │     VISITA          │
│      0       │       0             │
├─────────────────────────────────────┤
│              RELOJ                  │
│           20:00.00                  │
├─────────────────────────────────────┤
│            DESCANSO                 │
│           05:00.00                  │
├─────────────────────────────────────┤
│         CONTROL PANEL               │
│  [START] [STOP] [RESET]             │
│  [INICIAR DESCANSO] [TERMINAR]      │
│  [RESET DESCANSO]                   │
│  [+2] [+3] [-1]  [+2] [+3] [-1]    │
└─────────────────────────────────────┘
```

## 🚀 Ventajas de la Implementación

### ✅ **Separación de Responsabilidades**

- **Reloj principal**: Maneja el tiempo del juego
- **Reloj de descanso**: Maneja los descansos
- **Independencia**: Cada reloj funciona por separado

### ✅ **Automatización**

- **Transición automática** del descanso al juego
- **Callback automático** cuando termina el descanso
- **Sincronización perfecta** entre relojes

### ✅ **Control Manual**

- **Usuario puede** terminar el descanso manualmente
- **Flexibilidad** en el manejo de tiempos
- **Intervención humana** cuando sea necesario

### ✅ **Data Binding Automático**

- **Actualización automática** de displays
- **Sincronización perfecta** entre ViewModels y Views
- **Sin código boilerplate** para actualizaciones

## 🛠️ Uso Práctico

### **Escenario 1: Descanso Automático**

```go
// El usuario inicia un descanso
gameViewModel.StartRest()
// Automáticamente:
// 1. Reloj principal se detiene
// 2. Reloj de descanso inicia (1 minuto)
// 3. Cuando termina, reloj principal reinicia automáticamente
```

### **Escenario 2: Descanso Manual**

```go
// El usuario inicia un descanso
gameViewModel.StartRest()
// El usuario termina el descanso antes de tiempo
gameViewModel.StopRest()
// Automáticamente:
// 1. Reloj de descanso se detiene
// 2. Reloj principal inicia inmediatamente
```

## 🎯 Configuración

### **Tiempo de Descanso por Defecto:**

```go
var DefaultGameConfig = GameConfig{
    Quarters:         2,
    TimePerQuarter:   20 * time.Minute,
    RestTime:         5 * time.Minute,  // ← Tiempo de descanso
    TimePerTimeOut:   1 * time.Minute,
    TimePerExtraTime: 3 * time.Minute,
}
```

### **Personalización:**

```go
config := GameConfig{
    RestTime: 2 * time.Minute, // Descanso de 2 minutos
}
gameViewModel := NewGameViewModelWithConfig(config)
```

## 🎉 Resultado Final

La implementación del reloj de descanso proporciona:

- ✅ **Funcionalidad completa** de descansos
- ✅ **Automatización inteligente** de transiciones
- ✅ **Control manual** cuando sea necesario
- ✅ **Interfaz clara** y fácil de usar
- ✅ **Arquitectura MVVM** mantenida
- ✅ **Data binding automático** funcionando

**¡El reloj de descanso está completamente integrado y funcionando!** 🕐
