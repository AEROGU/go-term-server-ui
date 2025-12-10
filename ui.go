package main

import (
	"github.com/rivo/tview"
)

// TODO: Lo que ahora se ejecuta en main.go lo convertiremos en una librería apropiada aquí.

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
)

func InitializeUI() {
	App = tview.NewApplication().EnableMouse(true)

	serverInfo = tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetRegions(true) // Se establece un SetChangedFunc más adelante
	serverInfo.SetBorder(true).SetTitle("Server Info")

	logOutput = tview.NewTextView().SetDynamicColors(true).SetScrollable(true).SetMaxLines(1000).SetChangedFunc(func() { App.Draw() })
	logOutput.SetBorder(true).SetTitle("Server Logs")

	connectLogWriterToTextViewAndLog(logOutput)
	SetServerInfo("IP: [yellow]0.0.0.0[-]\nPort: [yellow]0[-]\nStatus: [red::l]Offline[-::L]") // Default info al iniciar.
	controlsFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	controlsFlex.SetBorder(true).SetTitle("Controls")

	// Aquí puedes agregar botones u otros controles a 'controls'

	stopServerButton := tview.NewButton("Stop Server").SetSelectedFunc(func() {
		// Lógica para detener el servidor
		SetServerInfo("IP: [yellow]127.0.0.1[-]\nPort: [yellow]3000[-]\nStatus: [gray]Offline[-]")
	})
	controlsFlex.AddItem(stopServerButton, 0, 1, false)

	lateralLeft := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(serverInfo, serverInfoWidth(), 1, false). // Top Left
		AddItem(controlsFlex, 3, 0, false)                // Bottom Left
	// Actualizar tamaño automáticamente cuando cambie el texto
	serverInfo.SetChangedFunc(func() {
		newHeight := serverInfoWidth()
		lateralLeft.ResizeItem(serverInfo, newHeight, 1)
		App.Draw()
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
	if err := App.Run(); err != nil {
		return err
	}
	isAppRunning = false
	return nil
}
