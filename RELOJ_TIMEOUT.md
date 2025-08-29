# ⏱️ Reloj de TimeOut - Funcionalidad Implementada

## 🎯 Descripción

Se ha implementado un **reloj de TimeOut separado** que funciona de manera independiente al reloj principal del juego. Este reloj cuenta 1 minuto y automáticamente reinicia el reloj principal cuando termina.

## 🏗️ Arquitectura Implementada

### 1. **TimeOutClockViewModel** (`src/viewmodels/timeout_clock_viewmodel.go`)

```go
type TimeOutClockViewModel struct {
    TimeLeft         *observable.ObservableDuration
    IsRunning        *observable.ObservableBool
    IsStopped        *observable.ObservableBool
    TimeOutTime      time.Duration
    ticker           *time.Ticker
    stopChan         chan bool
    onTimeOutFinished func() // Callback cuando termina el TimeOut
}
```

**Características:**

- ✅ **Reloj independiente** con su propio ticker
- ✅ **Data binding automático** con observables
- ✅ **Callback automático** cuando termina el TimeOut
- ✅ **Thread-safe** con canales de comunicación

### 2. **TimeOutClockView** (`src/views/timeout_clock_view.go`)

```go
type TimeOutClockView struct {
    ViewModel      *viewmodels.TimeOutClockViewModel
    Container      *fyne.Container
    TimeLabel      *canvas.Text
    ControlButtons *fyne.Container
}
```

**Características:**

- ✅ **Vista separada** del reloj de TimeOut
- ✅ **Data binding automático** con el ViewModel
- ✅ **Botones de control** específicos
- ✅ **Actualización automática** del display

## 🔄 Flujo de Funcionamiento

### **Secuencia Normal:**

1. **Reloj principal** está corriendo
2. **Usuario** hace clic en "INICIAR TIMEOUT"
3. **Reloj principal** se detiene automáticamente
4. **Reloj de TimeOut** inicia la cuenta regresiva (1 minuto)
5. **Cuando termina** el TimeOut:
   - Se ejecuta el callback automáticamente
   - **Reloj principal** se reinicia
   - **Reloj principal** inicia automáticamente
6. **Reloj de TimeOut** se resetea

### **Control Manual:**

- **Usuario** puede hacer clic en "TERMINAR TIMEOUT"
- **Reloj de TimeOut** se detiene
- **Reloj principal** inicia inmediatamente

## 🎮 Controles Disponibles

### **Reloj Principal:**

- **START**: Inicia el reloj principal
- **STOP**: Detiene el reloj principal
- **RESET**: Reinicia el reloj principal
- **INICIAR DESCANSO**: Detiene el principal e inicia el descanso
- **INICIAR TIMEOUT**: Detiene el principal e inicia el TimeOut

### **Reloj de TimeOut:**

- **INICIAR TIMEOUT**: Inicia el reloj de TimeOut
- **TERMINAR TIMEOUT**: Detiene el TimeOut e inicia el principal
- **RESET**: Reinicia el reloj de TimeOut

## 📊 Integración con GameViewModel

### **Nuevos Métodos Agregados:**

```go
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
```

### **Configuración Automática:**

```go
// En NewGameViewModel()
gameVM.TimeOutClock.SetOnTimeOutFinished(func() {
    // Cuando termina el TimeOut, reiniciar el reloj principal
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
│            TIMEOUT                  │
│           01:00.00                  │
├─────────────────────────────────────┤
│         CONTROL PANEL               │
│  [START] [STOP] [RESET]             │
│  [INICIAR DESCANSO] [TERMINAR]      │
│  [INICIAR TIMEOUT] [TERMINAR]       │
│  [RESET DESCANSO] [RESET TIMEOUT]   │
│  [+2] [+3] [-1]  [+2] [+3] [-1]    │
└─────────────────────────────────────┘
```

## 🚀 Ventajas de la Implementación

### ✅ **Separación de Responsabilidades**

- **Reloj principal**: Maneja el tiempo del juego
- **Reloj de descanso**: Maneja los descansos
- **Reloj de TimeOut**: Maneja los timeouts
- **Independencia**: Cada reloj funciona por separado

### ✅ **Automatización**

- **Transición automática** del TimeOut al juego
- **Callback automático** cuando termina el TimeOut
- **Sincronización perfecta** entre relojes

### ✅ **Control Manual**

- **Usuario puede** terminar el TimeOut manualmente
- **Flexibilidad** en el manejo de tiempos
- **Intervención humana** cuando sea necesario

### ✅ **Data Binding Automático**

- **Actualización automática** de displays
- **Sincronización perfecta** entre ViewModels y Views
- **Sin código boilerplate** para actualizaciones

## 🛠️ Uso Práctico

### **Escenario 1: TimeOut Automático**

```go
// El usuario inicia un TimeOut
gameViewModel.StartTimeOut()
// Automáticamente:
// 1. Reloj principal se detiene
// 2. Reloj de TimeOut inicia (1 minuto)
// 3. Cuando termina, reloj principal reinicia automáticamente
```

### **Escenario 2: TimeOut Manual**

```go
// El usuario inicia un TimeOut
gameViewModel.StartTimeOut()
// El usuario termina el TimeOut antes de tiempo
gameViewModel.StopTimeOut()
// Automáticamente:
// 1. Reloj de TimeOut se detiene
// 2. Reloj principal inicia inmediatamente
```

## 🎯 Configuración

### **Tiempo de TimeOut por Defecto:**

```go
var DefaultGameConfig = GameConfig{
    Quarters:         2,
    TimePerQuarter:   20 * time.Minute,
    RestTime:         5 * time.Minute,
    TimePerTimeOut:   1 * time.Minute,  // ← Tiempo de TimeOut
    TimePerExtraTime: 3 * time.Minute,
}
```

### **Personalización:**

```go
config := GameConfig{
    TimePerTimeOut: 30 * time.Second, // TimeOut de 30 segundos
}
gameViewModel := NewGameViewModelWithConfig(config)
```

## 🔄 Diferencias con el Reloj de Descanso

| Aspecto        | Reloj de Descanso      | Reloj de TimeOut    |
| -------------- | ---------------------- | ------------------- |
| **Duración**   | 5 minutos              | 1 minuto            |
| **Propósito**  | Descanso entre cuartos | Pausa táctica       |
| **Frecuencia** | Entre cuartos          | Durante el juego    |
| **Control**    | Manual y automático    | Manual y automático |

## 🎉 Resultado Final

La implementación del reloj de TimeOut proporciona:

- ✅ **Funcionalidad completa** de timeouts
- ✅ **Automatización inteligente** de transiciones
- ✅ **Control manual** cuando sea necesario
- ✅ **Interfaz clara** y fácil de usar
- ✅ **Arquitectura MVVM** mantenida
- ✅ **Data binding automático** funcionando

**¡El reloj de TimeOut está completamente integrado y funcionando!** ⏱️
