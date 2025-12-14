package main

import "strings"

// getLineCount returns the number of lines in the given text.
func getLineCount(text string) int {
	// Elimina saltos de línea finales para evitar contar una línea vacía extra
	text = strings.TrimRight(text, "\r\n")
	if len(text) == 0 {
		return 0
	}
	// El número de líneas es la cantidad de separadores '\n' + 1
	return strings.Count(text, "\n") + 1
}

// serverInfoWidth calculates the appropriate height for the serverInfo TextView
// based on its content, with a maximum height of 8 lines.
// The goal of having a maximum size is to avoid taking up too much space from
// the possible buttons in the controls section.
func serverInfoWidth() int {
	var serverInfoLineCount int = getLineCount(serverInfo.GetText(true))
	if serverInfoLineCount == 0 {
		return 0 // Si no hay texto, altura 0 y no tiene caso procesar lo demás.
	}
	if serverInfoLineCount > 8 {
		serverInfoLineCount = 8 // Altura máxima de 8 líneas, si se pasa deberá usar el scroll.
	}
	return serverInfoLineCount + 2 // +2 para los bordes
}

// SetServerInfo actualiza el texto del serverInfo TextView de manera segura
// dependiendo de si la aplicación está corriendo o no.
func SetServerInfo(text string) {
	if isAppRunning {
		App.QueueUpdateDraw(func() {
			serverInfo.SetText(text)
		})
	} else {
		serverInfo.SetText(text)
	}
}

// RefreshUI fuerza un redraw de la UI si la aplicación está corriendo.
// Útil para actualizar la interfaz después de cambios en segundo plano,
// como por ejemplo después de cambiar el texto o el color de fondo de
// un botón o un textview.
func RefreshUI() {
	if App == nil {
		return
	}

	if isAppRunning {
		// Desde gorutines en background: encolamos una actualización segura.
		App.QueueUpdateDraw(func() {})
		return
	}

	// Si la app no está corriendo (o estamos en el hilo de UI),
	// forzamos el dibujo inmediatamente.
	App.Draw()
}
