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
//
// Reglas:
//   - Antes de que la app corra (App == nil o !isAppRunning), se asume que
//     estamos en el hilo principal de inicialización y se puede llamar
//     directamente a SetText.
//   - Cuando la app está corriendo, la actualización se encola usando
//     App.QueueUpdateDraw **desde una goroutine separada**, ya que QueueUpdate*
//     está pensada para ser llamada desde gorutinas no-UI. Llamarla dentro de
//     callbacks de la propia UI puede provocar deadlocks.
func SetServerInfo(text string) {
	// Fase de inicialización: sin event loop todavía.
	if App == nil || !isAppRunning {
		if serverInfo != nil {
			serverInfo.SetText(text)
		}
		return
	}

	// App está corriendo: encolamos la actualización en el hilo de UI desde
	// una goroutine aparte para evitar bloquear el event loop.
	t := text
	go func() {
		App.QueueUpdateDraw(func() {
			serverInfo.SetText(t)
		})
	}()
}

// RefreshUI fuerza un redraw de la UI si la aplicación está corriendo.
// Útil para actualizar la interfaz después de cambios en segundo plano,
// como por ejemplo después de cambiar el texto o el color de fondo de
// un botón o un textview.
// Si la app no está corriendo, no hace nada (los cambios se renderizarán
// automáticamente cuando App.Run() inicie).
func RefreshUI() {
	if App == nil || !isAppRunning {
		return
	}

	// Encolamos una actualización segura desde una goroutine separada para
	// evitar llamar QueueUpdate*/Draw desde el propio event loop.
	go func() {
		App.QueueUpdateDraw(func() {})
	}()
}
