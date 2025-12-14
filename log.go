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
func (w *tuiLogWriter) Print(msg string) {
	w.Write([]byte(msg))
}

func (w *tuiLogWriter) Println(msg string) {
	w.Write([]byte(msg + "\n"))
}

// Printf agrega un mensaje formateado al log view de forma segura para la UI
func (w *tuiLogWriter) Printf(format string, a ...interface{}) {
	_, _ = w.Write([]byte(fmt.Sprintf(format, a...)))
}
