# Instrucciones de Uso - Basketball Score Counter

## 🏀 Descripción
Esta aplicación te permite contar puntos y cronometrar partidos de baloncesto de manera fácil y rápida.

## 🚀 Cómo ejecutar la aplicación

### Opción 1: Ejecutar directamente
```bash
go run main.go
```

### Opción 2: Compilar y ejecutar
```bash
go build -o BasketballCounter.exe main.go
./BasketballCounter.exe
```

### Opción 3: Usar el script de compilación
```bash
build.bat
```

## 📋 Funcionalidades principales

### 1. Configuración de equipos
- **Cambiar nombres**: Haz clic en los campos de texto y escribe los nombres de los equipos
- Los nombres se actualizan automáticamente en la interfaz

### 2. Contador de puntos
Para cada equipo tienes botones para:
- **+1**: Sumar 1 punto (tiro libre)
- **+2**: Sumar 2 puntos (canasta normal)
- **+3**: Sumar 3 puntos (triple)
- **-1**: Descontar 1 punto (para corregir errores)

### 3. Cronómetro
- **Iniciar/Pausar**: Controla el tiempo del cuarto actual
- **Reiniciar Tiempo**: Vuelve a poner el tiempo completo del cuarto actual
- **Siguiente Cuarto**: Avanza al siguiente cuarto (máximo 4 cuartos)

### 4. Configuración de tiempo
- Selecciona la duración de cada cuarto:
  - 8 minutos
  - 10 minutos (por defecto)
  - 12 minutos

### 5. Nuevo juego
- **Nuevo Juego**: Reinicia todo el partido (puntuación y tiempo)

## 🎯 Flujo de trabajo típico

1. **Antes del partido**:
   - Cambia los nombres de los equipos
   - Selecciona la duración de los cuartos
   - Verifica que todo esté en cero

2. **Durante el partido**:
   - Haz clic en "Iniciar" para comenzar el cronómetro
   - Usa los botones +1, +2, +3 para sumar puntos según corresponda
   - Si hay un error, usa el botón -1 para corregir
   - Pausa el cronómetro si hay interrupciones

3. **Al final de cada cuarto**:
   - El cronómetro se detendrá automáticamente
   - Aparecerá una notificación
   - Haz clic en "Siguiente Cuarto" para continuar

4. **Al final del partido**:
   - Aparecerá una notificación de "Fin del Juego"
   - Puedes iniciar un nuevo juego con "Nuevo Juego"

## ⌨️ Atajos de teclado (futuras versiones)
- Espacio: Iniciar/Pausar cronómetro
- R: Reiniciar tiempo
- N: Nuevo juego
- 1, 2, 3: Sumar puntos al equipo local
- Q, W, E: Sumar puntos al equipo visitante

## 🔧 Solución de problemas

### La aplicación no se ejecuta
1. Verifica que tienes Go instalado: `go version`
2. Instala las dependencias: `go mod tidy`
3. Ejecuta: `go run main.go`

### El cronómetro no funciona correctamente
1. Asegúrate de que el cronómetro esté iniciado
2. Si está pausado, haz clic en "Iniciar"
3. Si hay problemas, usa "Reiniciar Tiempo"

### Los puntos no se suman
1. Verifica que estés haciendo clic en los botones correctos
2. Cada equipo tiene sus propios botones (izquierda y derecha)
3. Usa el botón -1 si necesitas corregir un error

## 📱 Características técnicas

- **Interfaz nativa**: Se integra perfectamente con Windows
- **Responsive**: Se adapta al tamaño de la ventana
- **Thread-safe**: Manejo seguro del cronómetro
- **Memoria eficiente**: Uso mínimo de recursos

## 🆘 Soporte

Si tienes problemas:
1. Verifica que tienes la versión correcta de Go (1.22.1+)
2. Asegúrate de que todas las dependencias estén instaladas
3. Revisa que no haya otros programas usando los mismos recursos

## 🔄 Próximas versiones

Funcionalidades planeadas:
- Guardar/cargar partidos
- Estadísticas detalladas
- Sonidos personalizables
- Modo oscuro
- Exportar resultados
- Soporte para múltiples idiomas 