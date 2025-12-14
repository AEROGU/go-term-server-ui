package main

import (
	"fmt"

	"github.com/rivo/tview"
)

// LogWriter es el escritor personalizado que envía los logs al TextView, implementa io.Writer
// Se inicializa en InitializeUI()
var LogWriter *tuiLogWriter

// tuiLogWriter implementa io.Writer para enviar logs al TextView
type tuiLogWriter struct {
	view *tview.TextView
}

// Write escribe los datos en el TextView de forma segura para la UI
// No guarda nada en algún log, solo lo muestra directamente
// Si quieres que muestre los mensajed de log puedes redirigir
// la salida de un log así: "log.SetOutput(LogWriter)" para mostrarlo en la interfaz
// o redirigirlo a mas de 1 lugar si quieres también guardarlo en un archivo aparte así:
// "log.SetOutput(io.MultiWriter(myFileWriter, LogWriter))"
func (w *tuiLogWriter) Write(p []byte) (n int, err error) {
	// Es crucial usar QueueUpdateDraw porque Write puede ser llamado desde
	// cualquier goroutine, y tview no es thread-safe por defecto.
	App.QueueUpdateDraw(func() {
		w.view.Write(p)
		w.view.ScrollToEnd() // Mantener el scroll abajo
	})
	return len(p), nil
}

// Print agrega un mensaje al log view de forma segura para la UI
// No guarda nada en algún log, solo lo muestra directamente
func (w *tuiLogWriter) Print(msg string) {
	w.Write([]byte(msg))
}

// Println agrega un mensaje con salto de línea al log view de forma segura para la UI
// No guarda nada en algún log, solo lo muestra directamente
func (w *tuiLogWriter) Println(msg string) {
	w.Write([]byte(msg + "\n"))
}

// Printf agrega un mensaje formateado al log view de forma segura para la UI
// No guarda nada en algún log, solo lo muestra directamente
func (w *tuiLogWriter) Printf(format string, a ...interface{}) {
	_, _ = w.Write([]byte(fmt.Sprintf(format, a...)))
}
