# Estructura Refactorizada - Patrón MVC

## Descripción General

El proyecto ha sido refactorizado siguiendo el patrón **Model-View-Controller (MVC)** para separar claramente las responsabilidades entre la lógica de negocio y la interfaz de usuario.

## Estructura de Directorios

```
src/
├── controllers/          # Controladores (Coordinación)
│   └── game_controller.go
├── guiModels/           # Vistas (Interfaz de Usuario)
│   ├── team.go
│   ├── clock.go
│   └── guiutils.go
├── models/              # Modelos (Lógica de Negocio)
│   ├── game.go
│   └── config.go
├── mainTab.go           # Configuración de la interfaz principal
└── backup.go
```

## Componentes

### 1. Modelos (Lógica de Negocio)

#### `src/models/game.go`

- **Team**: Estructura pura con lógica de negocio del equipo
- **Clock**: Lógica del reloj del juego
- **Game**: Modelo principal del juego
- **GameConfig**: Configuración del juego

**Características:**

- ✅ Sin dependencias de GUI
- ✅ Lógica de negocio pura
- ✅ Fácil de testear
- ✅ Reutilizable

### 2. Vistas (Interfaz de Usuario)

#### `src/guiModels/team.go`

- **TeamGuiModel**: Interfaz visual del equipo
- **Team**: Modelo de datos local (evita ciclos de importación)

#### `src/guiModels/clock.go`

- **ClockGuiModel**: Interfaz visual del reloj
- **Clock**: Modelo de datos local del reloj

**Características:**

- ✅ Solo se encarga de la presentación
- ✅ Delega lógica al modelo
- ✅ Actualiza automáticamente la UI

### 3. Controladores (Coordinación)

#### `src/controllers/game_controller.go`

- **GameController**: Coordina entre modelos y vistas
- **GameConfig**: Configuración centralizada

**Características:**

- ✅ Orquesta la interacción entre modelos y vistas
- ✅ Maneja la lógica de aplicación
- ✅ Punto único de entrada para operaciones del juego

## Flujo de Datos

```
Usuario → Controlador → Modelo → Vista → Usuario
```

### Ejemplo: Añadir puntos

1. **Usuario** hace clic en botón "+2"
2. **Controlador** recibe la acción
3. **Modelo** actualiza el puntaje
4. **Vista** se actualiza automáticamente
5. **Usuario** ve el cambio en pantalla

## Ventajas de la Nueva Estructura

### ✅ Separación de Responsabilidades

- **Modelos**: Solo lógica de negocio
- **Vistas**: Solo presentación
- **Controladores**: Solo coordinación

### ✅ Mantenibilidad

- Cambios en la UI no afectan la lógica
- Cambios en la lógica no afectan la UI
- Fácil de extender y modificar

### ✅ Testabilidad

- Modelos se pueden testear independientemente
- Vistas se pueden testear con mocks
- Controladores se pueden testear aisladamente

### ✅ Reutilización

- Modelos se pueden usar en diferentes interfaces
- Lógica de negocio es portable
- Fácil de migrar a otras tecnologías

## Uso

### Crear un nuevo juego:

```go
gameController := controllers.NewGameController()
```

### Acceder a componentes:

```go
homeTeam := gameController.GetHomeTeamGui()
clock := gameController.GetClockGui()
```

### Operaciones del juego:

```go
gameController.StartGame()
gameController.AddScoreToHomeTeam(2)
gameController.ResetGame()
```

## Migración

La refactorización mantiene la funcionalidad existente pero con mejor estructura:

- ✅ Todas las funciones originales siguen funcionando
- ✅ Interfaz de usuario idéntica
- ✅ Código más organizado y mantenible
- ✅ Preparado para futuras extensiones
