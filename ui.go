package main

import "github.com/rivo/tview"

// TODO: Lo que ahora se ejecuta en main.go lo convertiremos en una librería apropiada aquí.

type UI struct {
	// Aquí irán los elementos de la UI y métodos relacionados
	isAppRunning bool
	elements     widgets
	LogWriter    *LogWriter
}

type widgets struct {
	// Aquí irán los widgets de la UI
	App        *tview.Application
	serverInfo *tview.TextView
	logOutput  *tview.TextView
}

// TODO: Mover logWriter aquí también y sus métodos relacionados.
