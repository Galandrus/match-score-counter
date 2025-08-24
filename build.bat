@echo off
echo Compilando Basketball Score Counter...
echo.

REM Verificar que Go esté instalado
go version >nul 2>&1
if errorlevel 1 (
    echo ERROR: Go no está instalado o no está en el PATH
    echo Por favor instala Go desde https://golang.org/dl/
    pause
    exit /b 1
)

REM Instalar dependencias
echo Instalando dependencias...
go mod tidy
if errorlevel 1 (
    echo ERROR: No se pudieron instalar las dependencias
    pause
    exit /b 1
)

REM Compilar la aplicación
echo Compilando aplicación...
go build -ldflags="-s -w" -o BasketballCounter.exe main.go
if errorlevel 1 (
    echo ERROR: Error durante la compilación
    pause
    exit /b 1
)

echo.
echo ¡Compilación exitosa!
echo El ejecutable BasketballCounter.exe ha sido creado
echo.
echo Para ejecutar la aplicación, haz doble clic en BasketballCounter.exe
echo.
pause