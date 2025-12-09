package main

import (
	"log"
	"time"

	"github.com/rivo/tview"
)

var App = tview.NewApplication().EnableMouse(true)
var isAppRunning = false                                                                         // Bandera para controlar el estado de la app
var serverInfo = tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetRegions(true) // Se establece un SetChangedFunc más adelante
var logOutput = tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetMaxLines(1000).SetChangedFunc(func() { App.Draw() })

// CONECTAR EL LOGGER
// Creamos nuestro writer personalizado
var logWriter = &LogWriter{view: logOutput}

func main() {

	logOutput.SetBorder(true).SetTitle("Server Logs")

	serverInfo.SetBorder(true).SetTitle("Server Info")
	setServerInfo("IP: [yellow]127.0.0.1[-]\nPort: [yellow]3000[-]\nStatus: [green::l]Online[-]\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10")

	controls := tview.NewFlex().SetDirection(tview.FlexRow)
	controls.SetBorder(true).SetTitle("Controls")
	// Aquí puedes agregar botones u otros controles a 'controls'

	stopServerButton := tview.NewButton("Stop Server").SetSelectedFunc(func() {
		// Lógica para detener el servidor
		setServerInfo("IP: [yellow]127.0.0.1[-]\nPort: [yellow]3000[-]\nStatus: [gray]Offline[-]")
	})
	controls.AddItem(stopServerButton, 3, 0, false)

	lateralLeft := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(serverInfo, serverInfoWidth(), 1, false).                         // Top Left
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Controls"), 0, 1, false) // Bottom Left

	// Actualizar tamaño automáticamente cuando cambie el texto
	serverInfo.SetChangedFunc(func() {
		newHeight := serverInfoWidth()
		lateralLeft.ResizeItem(serverInfo, newHeight, 1)
		App.Draw()
	})

	mainFlex := tview.NewFlex().
		AddItem(lateralLeft, 0, 1, false).
		AddItem(logOutput, 0, 2, false)

	// Redirigimos el log global de Go a nuestro TextView ---------------------------
	log.SetOutput(logWriter)
	// Opcional: Quitar flags de fecha si quieres controlarlo tú
	// log.SetFlags(0)
	// ------------------------------------------------------------------------------

	// logWriter.Write([]byte("[green]Application started. Logs will appear here...[-]\n"))

	// Test: Actualizar info del servidor después de 5 segundos -----------------
	time.AfterFunc(5*time.Second, func() {
		setServerInfo("Test")
	})

	App.SetRoot(mainFlex, true)
	isAppRunning = true
	if err := App.Run(); err != nil {
		panic(err)
	}
	isAppRunning = false
}

// LogWriter implementa io.Writer para enviar logs al TextView
type LogWriter struct {
	view *tview.TextView
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	// Es crucial usar QueueUpdateDraw porque Write puede ser llamado desde
	// cualquier goroutine, y tview no es thread-safe por defecto.
	App.QueueUpdateDraw(func() {
		w.view.Write(p)
		w.view.ScrollToEnd() // Mantener el scroll abajo
	})
	return len(p), nil
}
