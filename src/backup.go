package src

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// )

// type BasketballGame struct {
// 	homeScore        int
// 	awayScore        int
// 	quarter          int
// 	timeLeft         time.Duration
// 	isRunning        bool
// 	timer            *time.Timer
// 	quarterTime      time.Duration
// 	homeTeamName     string
// 	awayTeamName     string
// 	updateTimeFunc   func()
// 	updateScoresFunc func()
// 	homeScoreLabel   *widget.Label
// 	awayScoreLabel   *widget.Label
// }

// func main() {
// 	fmt.Println("start app")
// 	myApp := app.New()
// 	myApp.SetIcon(theme.ColorChromaticIcon())
// 	window := myApp.NewWindow("Cestoball Score Counter")
// 	window.Resize(fyne.NewSize(1000, 700))

// 	game := &BasketballGame{
// 		quarter:      1,
// 		quarterTime:  10 * time.Minute,
// 		timeLeft:     10 * time.Minute,
// 		homeTeamName: "Local",
// 		awayTeamName: "Visita",
// 	}

// 	// Crear pestañas
// 	tabs := container.NewAppTabs()

// 	// Pestaña 1: Marcador Principal
// 	mainScoreTab := createMainScoreTab(game)
// 	tabs.Append(container.NewTabItem("Marcador", mainScoreTab))

// 	// Pestaña 2: Configuración
// 	configTab := createConfigTab(game)
// 	tabs.Append(container.NewTabItem("Configuración", configTab))

// 	window.SetContent(tabs)
// 	window.ShowAndRun()
// 	fmt.Println("end app")
// }

// func createMainScoreTab(game *BasketballGame) fyne.CanvasObject {
// 	// Área principal de visualización (izquierda)
// 	mainDisplay := createMainDisplay(game)

// 	// Panel de control (derecha)
// 	controlPanel := createControlPanel(game)

// 	// Layout horizontal: display principal + panel de control
// 	return container.NewHSplit(mainDisplay, controlPanel)
// }

// func createMainDisplay(game *BasketballGame) fyne.CanvasObject {
// 	// Título principal
// 	title := canvas.NewText("MARCADOR", nil)
// 	title.TextSize = 24
// 	title.TextStyle = fyne.TextStyle{Bold: true}
// 	title.Alignment = fyne.TextAlignCenter

// 	// Marcadores grandes
// 	homeScoreDisplay := createLargeScoreDisplay("LOCAL", game.homeScore)
// 	awayScoreDisplay := createLargeScoreDisplay("VISITA", game.awayScore)

// 	// Área de marcadores
// 	scoresArea := container.NewHBox(homeScoreDisplay, awayScoreDisplay)

// 	// Reloj principal
// 	clockLabel := canvas.NewText("RELOJ", nil)
// 	title.TextSize = 24

// 	clockLabel.TextStyle = fyne.TextStyle{Bold: true}
// 	clockLabel.Alignment = fyne.TextAlignCenter

// 	timeDisplay := canvas.NewText("00:00", nil)
// 	title.TextSize = 32
// 	timeDisplay.TextStyle = fyne.TextStyle{Bold: true}
// 	timeDisplay.Alignment = fyne.TextAlignCenter

// 	// Actualizar el tiempo
// 	updateTime := func() {
// 		timeDisplay.Text = formatTime(game.timeLeft)
// 		timeDisplay.Refresh()
// 	}

// 	// Función para actualizar marcadores
// 	updateScores := func() {
// 		// Por ahora, solo actualizamos el tiempo
// 		// Los marcadores se actualizarán cuando se presionen los botones
// 	}

// 	// Guardar referencias para actualizaciones
// 	game.updateTimeFunc = updateTime
// 	game.updateScoresFunc = updateScores

// 	clockArea := container.NewVBox(clockLabel, timeDisplay)

// 	// Layout principal del display
// 	mainContent := container.NewVBox(
// 		title,
// 		scoresArea,
// 		clockArea,
// 	)

// 	return container.NewPadded(mainContent)
// }

// func createLargeScoreDisplay(team string, score int) fyne.CanvasObject {
// 	teamLabel := widget.NewLabel(team)
// 	teamLabel.TextStyle = fyne.TextStyle{Bold: true}
// 	teamLabel.Alignment = fyne.TextAlignCenter

// 	scoreLabel := widget.NewLabel(strconv.Itoa(score))
// 	scoreLabel.TextStyle = fyne.TextStyle{Bold: true}
// 	scoreLabel.Alignment = fyne.TextAlignCenter

// 	return container.NewVBox(teamLabel, layout.NewSpacer(), scoreLabel)
// }

// func createControlPanel(game *BasketballGame) fyne.CanvasObject {
// 	// Cronómetro
// 	clockLabel := widget.NewLabel("Cronómetro:")
// 	clockLabel.TextStyle = fyne.TextStyle{Bold: true}

// 	timeLabel := widget.NewLabel("00:00")
// 	timeLabel.TextStyle = fyne.TextStyle{Bold: true}
// 	timeLabel.Alignment = fyne.TextAlignCenter

// 	// Botones de control del tiempo
// 	timeControls := container.NewHBox(
// 		widget.NewButton("+", func() {
// 			game.timeLeft += time.Minute
// 			if game.updateTimeFunc != nil {
// 				game.updateTimeFunc()
// 			}
// 		}),
// 		widget.NewButton("-", func() {
// 			if game.timeLeft > time.Minute {
// 				game.timeLeft -= time.Minute
// 				if game.updateTimeFunc != nil {
// 					game.updateTimeFunc()
// 				}
// 			}
// 		}),
// 		timeLabel,
// 		widget.NewButton("+", func() {
// 			game.timeLeft += 10 * time.Second
// 			if game.updateTimeFunc != nil {
// 				game.updateTimeFunc()
// 			}
// 		}),
// 		widget.NewButton("-", func() {
// 			if game.timeLeft > 10*time.Second {
// 				game.timeLeft -= 10 * time.Second
// 				if game.updateTimeFunc != nil {
// 					game.updateTimeFunc()
// 				}
// 			}
// 		}),
// 	)

// 	// Botones de control del cronómetro
// 	var startStopButton *widget.Button
// 	startStopButton = widget.NewButton("Iniciar", func() {
// 		if !game.isRunning {
// 			game.isRunning = true
// 			startStopButton.SetText("Pausar")
// 			game.timer = time.NewTimer(1 * time.Second)
// 			go func() {
// 				for range game.timer.C {
// 					if game.isRunning {
// 						game.timeLeft -= time.Second
// 						if game.timeLeft <= 0 {
// 							game.timeLeft = 0
// 							game.isRunning = false
// 							startStopButton.SetText("Iniciar")
// 							dialog.ShowInformation("Fin de Cuarto",
// 								fmt.Sprintf("¡Fin del cuarto %d!", game.quarter), nil)
// 						}
// 						if game.updateTimeFunc != nil {
// 							game.updateTimeFunc()
// 						}
// 					}
// 					if game.isRunning {
// 						game.timer.Reset(1 * time.Second)
// 					}
// 				}
// 			}()
// 		} else {
// 			game.isRunning = false
// 			startStopButton.SetText("Iniciar")
// 			if game.timer != nil {
// 				game.timer.Stop()
// 			}
// 		}
// 	})

// 	resetButton := widget.NewButton("Reiniciar", func() {
// 		game.timeLeft = game.quarterTime
// 		if game.updateTimeFunc != nil {
// 			game.updateTimeFunc()
// 		}
// 		game.isRunning = false
// 		startStopButton.SetText("Iniciar")
// 		if game.timer != nil {
// 			game.timer.Stop()
// 		}
// 	})

// 	// Equipo Local
// 	localLabel := widget.NewLabel("Local")
// 	localLabel.TextStyle = fyne.TextStyle{Bold: true}
// 	localLabel.Alignment = fyne.TextAlignCenter

// 	dobleLocalBtn := widget.NewButton("Doble", func() {
// 		game.homeScore += 2
// 		if game.updateScoresFunc != nil {
// 			game.updateScoresFunc()
// 		}
// 	})

// 	tripleLocalBtn := widget.NewButton("Triple", func() {
// 		game.homeScore += 3
// 		if game.updateScoresFunc != nil {
// 			game.updateScoresFunc()
// 		}
// 	})

// 	localTantos := widget.NewLabel("TANTOS")
// 	localTantos.TextStyle = fyne.TextStyle{Bold: true}
// 	localTantos.Alignment = fyne.TextAlignCenter

// 	// Equipo Visitante
// 	visitaLabel := widget.NewLabel("Visita")
// 	visitaLabel.TextStyle = fyne.TextStyle{Bold: true}
// 	visitaLabel.Alignment = fyne.TextAlignCenter

// 	dobleVisitaBtn := widget.NewButton("Doble", func() {
// 		game.awayScore += 2
// 		if game.updateScoresFunc != nil {
// 			game.updateScoresFunc()
// 		}
// 	})

// 	tripleVisitaBtn := widget.NewButton("Triple", func() {
// 		game.awayScore += 3
// 		if game.updateScoresFunc != nil {
// 			game.updateScoresFunc()
// 		}
// 	})

// 	visitaTantos := widget.NewLabel("TANTOS")
// 	visitaTantos.TextStyle = fyne.TextStyle{Bold: true}
// 	visitaTantos.Alignment = fyne.TextAlignCenter

// 	// Layout de equipos
// 	teamsLayout := container.NewGridWithColumns(2,
// 		container.NewVBox(localLabel, dobleLocalBtn, tripleLocalBtn, localTantos),
// 		container.NewVBox(visitaLabel, dobleVisitaBtn, tripleVisitaBtn, visitaTantos),
// 	)

// 	// Layout principal del panel de control
// 	controlContent := container.NewVBox(
// 		clockLabel,
// 		timeControls,
// 		container.NewHBox(startStopButton, resetButton),
// 		widget.NewSeparator(),
// 		teamsLayout,
// 	)

// 	return container.NewPadded(controlContent)
// }

// func createConfigTab(game *BasketballGame) fyne.CanvasObject {
// 	// Configuración del tiempo por cuarto
// 	timeLabel := widget.NewLabel("Tiempo por cuarto:")
// 	timeOptions := []string{"8 minutos", "10 minutos", "12 minutos"}
// 	timeSelect := widget.NewSelect(timeOptions, func(selected string) {
// 		switch selected {
// 		case "8 minutos":
// 			game.quarterTime = 8 * time.Minute
// 		case "10 minutos":
// 			game.quarterTime = 10 * time.Minute
// 		case "12 minutos":
// 			game.quarterTime = 12 * time.Minute
// 		}
// 		if !game.isRunning {
// 			game.timeLeft = game.quarterTime
// 			if game.updateTimeFunc != nil {
// 				game.updateTimeFunc()
// 			}
// 		}
// 	})
// 	timeSelect.SetSelected("10 minutos")

// 	// Nombres de equipos
// 	homeTeamLabel := widget.NewLabel("Nombre Equipo Local:")
// 	homeTeamEntry := widget.NewEntry()
// 	homeTeamEntry.SetText(game.homeTeamName)
// 	homeTeamEntry.OnChanged = func(name string) {
// 		game.homeTeamName = name
// 	}

// 	awayTeamLabel := widget.NewLabel("Nombre Equipo Visitante:")
// 	awayTeamEntry := widget.NewEntry()
// 	awayTeamEntry.SetText(game.awayTeamName)
// 	awayTeamEntry.OnChanged = func(name string) {
// 		game.awayTeamName = name
// 	}

// 	// Botón para reiniciar todo el juego
// 	resetGameButton := widget.NewButton("Nuevo Juego", func() {
// 		dialog.ShowConfirm("Nuevo Juego", "¿Estás seguro de que quieres reiniciar todo el juego?", func(confirm bool) {
// 			if confirm {
// 				game.homeScore = 0
// 				game.awayScore = 0
// 				game.quarter = 1
// 				game.timeLeft = game.quarterTime
// 				game.isRunning = false

// 				if game.updateScoresFunc != nil {
// 					game.updateScoresFunc()
// 				}
// 				if game.updateTimeFunc != nil {
// 					game.updateTimeFunc()
// 				}

// 				if game.timer != nil {
// 					game.timer.Stop()
// 				}
// 			}
// 		}, nil)
// 	})

// 	// Layout de configuración
// 	configContent := container.NewVBox(
// 		timeLabel,
// 		timeSelect,
// 		widget.NewSeparator(),
// 		homeTeamLabel,
// 		homeTeamEntry,
// 		awayTeamLabel,
// 		awayTeamEntry,
// 		widget.NewSeparator(),
// 		resetGameButton,
// 	)

// 	return container.NewPadded(configContent)
// }

// func formatTime(d time.Duration) string {
// 	minutes := int(d.Minutes())
// 	seconds := int(d.Seconds()) % 60
// 	return fmt.Sprintf("%02d:%02d", minutes, seconds)
// }
