# Basketball Score Counter

Una aplicación Windows simple para contar puntos y cronometrar partidos de baloncesto.

## Características

- **Contador de puntos**: Botones para sumar 1, 2 o 3 puntos a cada equipo
- **Cronómetro**: Temporizador configurable por cuarto (8, 10 o 12 minutos)
- **Control de cuartos**: Avance automático entre cuartos
- **Nombres de equipos**: Campos editables para personalizar nombres
- **Interfaz moderna**: Diseño limpio y fácil de usar
- **Notificaciones**: Alertas al final de cada cuarto

## Requisitos

- Go 1.22.1 o superior
- Windows 10/11

## Instalación

1. **Instalar Go** (si no lo tienes):

   - Descarga Go desde [golang.org](https://golang.org/dl/)
   - Instala siguiendo las instrucciones oficiales

2. **Clonar o descargar el proyecto**:

   ```bash
   git clone <tu-repositorio>
   cd cestoballCounter
   ```

3. **Instalar dependencias**:

   ```bash
   go mod tidy
   ```

4. **Compilar y ejecutar**:
   ```bash
   go run main.go
   ```

## Uso

### Configuración inicial

1. **Nombres de equipos**: Edita los campos de texto para cambiar los nombres de los equipos
2. **Tiempo por cuarto**: Selecciona la duración deseada (8, 10 o 12 minutos)

### Durante el partido

1. **Iniciar cronómetro**: Haz clic en "Iniciar" para comenzar el tiempo
2. **Pausar**: Haz clic en "Pausar" para detener el cronómetro
3. **Sumar puntos**: Usa los botones +1, +2, +3 para cada equipo
4. **Descontar puntos**: Usa el botón -1 si necesitas corregir un error
5. **Siguiente cuarto**: Haz clic en "Siguiente Cuarto" cuando termine un período

### Controles adicionales

- **Reiniciar Tiempo**: Vuelve a poner el tiempo del cuarto actual
- **Nuevo Juego**: Reinicia todo el partido (puntuación y tiempo)

## Estructura del proyecto

```
cestoballCounter/
├── main.go          # Archivo principal de la aplicación
├── go.mod           # Dependencias del proyecto
└── README.md        # Este archivo
```

## Tecnologías utilizadas

- **Go**: Lenguaje de programación principal
- **Fyne**: Framework GUI multiplataforma
- **Go Modules**: Gestión de dependencias

## Compilación para distribución

Para crear un ejecutable independiente:

```bash
go build -o BasketballCounter.exe main.go
```

El archivo `BasketballCounter.exe` se puede ejecutar en cualquier Windows sin necesidad de instalar Go.

## Características técnicas

- **Interfaz nativa**: Se integra perfectamente con Windows
- **Responsive**: Se adapta al tamaño de la ventana
- **Thread-safe**: Manejo seguro de concurrencia para el cronómetro
- **Memoria eficiente**: Uso mínimo de recursos del sistema

## Solución de problemas

### Error de dependencias

Si obtienes errores de dependencias:

```bash
go mod download
go mod tidy
```

### Error de compilación

Asegúrate de tener Go 1.22.1 o superior:

```bash
go version
```

### La aplicación no se ejecuta

Verifica que tienes permisos de administrador si es necesario.

## Contribuir

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature
3. Commit tus cambios
4. Push a la rama
5. Abre un Pull Request

## Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo LICENSE para más detalles.
