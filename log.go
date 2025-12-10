package main

import (
	"log"

	"github.com/rivo/tview"
)

// Creamos nuestro writer que implementa io.Writter
var logWriter *LogWriter

// connectLogWriterToTextViewAndLog conecta el log global de Go a un TextView específico
func connectLogWriterToTextViewAndLog(view *tview.TextView) {
	logWriter = &LogWriter{view: view}
	log.SetOutput(logWriter)
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
