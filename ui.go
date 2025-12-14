package main

import (
	"github.com/rivo/tview"
)

// Variables globales de la aplicación.
// Se eligió hacerlas globales en lugar de usar structs porque solo habrá una instancia
// de la aplicación en todo momento, y esto simplifica el acceso desde distintas funciones.
var (
	// App es la aplicación tview principal
	App             *tview.Application
	isAppRunning    bool = false
	isUIInitialized bool = false
	// Aquí se muestra información sobre el servidor, se establece un SetChangedFunc más adelante para actualizar su tamaño dinámicamente
	serverInfo *tview.TextView
	logOutput  *tview.TextView
	// controlsFlex contendrá los botones y otros controles
	controlsFlex *tview.Flex
)

func InitializeUI() {
	App = tview.NewApplication().EnableMouse(true)

	serverInfo = tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetRegions(true) // Se establece un SetChangedFunc más adelante
	serverInfo.SetBorder(true).SetTitle("Server Info")

	logOutput = tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetMaxLines(1000)
	logOutput.SetBorder(true).SetTitle("Server Logs")

	LogWriter = &tuiLogWriter{view: logOutput} // Inicializar logWriter

	// SetServerInfo("IP: [yellow]0.0.0.0[-]\nPort: [yellow]0[-]\nStatus: [red::l]Offline[-::L]") // Default info al iniciar.
	controlsFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	controlsFlex.SetBorder(true).SetTitle("Controls")

	lateralLeft := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(serverInfo, serverInfoWidth(), 1, false). // Top Left
		AddItem(controlsFlex, 3, 0, false)                // Bottom Left
	// Actualizar tamaño automáticamente cuando cambie el texto.
	// Aquí no llamamos a RefreshUI porque este callback se ejecuta
	// dentro del hilo de UI cuando SetServerInfo actualiza el texto a
	// través de App.QueueUpdateDraw (cuando la app está corriendo) o de
	// forma directa antes de App.Run(). El redraw ya está garantizado.
	serverInfo.SetChangedFunc(func() {
		newHeight := serverInfoWidth()
		lateralLeft.ResizeItem(serverInfo, newHeight, 1)
	})

	mainFlex := tview.NewFlex().
		AddItem(lateralLeft, 0, 1, false).
		AddItem(logOutput, 0, 2, false)

	App.SetRoot(mainFlex, true)

	isUIInitialized = true
}

func RunUI() error {
	if !isUIInitialized {
		InitializeUI()
	}
	isAppRunning = true
	// connectLogWriterToTextViewAndLog(logOutput)
	// defer log.SetOutput(nil) // Restablece al valor predeterminado (stderr)
	if err := App.Run(); err != nil {
		return err
	}
	isAppRunning = false
	return nil
}
