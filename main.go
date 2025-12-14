package main

import (
	"log"
	"time"

	"github.com/AEROGU/go-term-server-ui/serverinfo"

	"github.com/rivo/tview"
)

func main() {
	InitializeUI() // Inicializar la UI antes de usarla o configurarla.

	// Setear info inicial del servidor
	srvInfo := serverinfo.Info{IP: "127.0.0.1", Port: 3000, Status: serverinfo.ServerStatusStoped}
	SetServerInfo(srvInfo.String())

	// Test: Actualizar info del servidor después de 5 segundos -----------------
	// time.AfterFunc(5*time.Second, func() {
	// 	SetServerInfo(srvInfo.String())
	// })

	// Agregar botón de ejemplo para detener/iniciar el servidor
	var btn *tview.Button
	btn = ControlsAddButton("Start server", func() {
		// Closure para capturar btn. Si no se captura, no se puede cambiar la etiqueta.
		func(b *tview.Button) {

			switch srvInfo.Status {
			case serverinfo.ServerStatusStoped:
				b.SetDisabled(true)
				srvInfo.Status = serverinfo.ServerStatusStarting
				time.AfterFunc(3*time.Second, func() { // Simular tiempo de arranque
					srvInfo.Status = serverinfo.ServerStatusRunning
					b.SetLabel("Stop server")
					b.SetDisabled(false)
					log.Println("[blue]Server started.[-]")
					SetServerInfo(srvInfo.String()) // Actualizar info del servidor. Esto también invoca App.QueueUpdateDraw, por lo que no es necesario llamar a RefreshUI()
				})
			case serverinfo.ServerStatusRunning:
				b.SetDisabled(true)
				srvInfo.Status = serverinfo.ServerStatusStopping
				time.AfterFunc(3*time.Second, func() { // Simular tiempo de parada
					srvInfo.Status = serverinfo.ServerStatusStoped
					b.SetLabel("Start server")
					b.SetDisabled(false)
					log.Printf("[blue]Server stopped.[-]")
					SetServerInfo(srvInfo.String()) // Actualizar info del servidor. Esto también invoca App.QueueUpdateDraw, por lo que no es necesario llamar a RefreshUI()
				})
			default:
				// No debe poderse clicar el botón en otros estados
				// pero por si acaso, simplemente imprimir en el log que el servidor está en transición
				log.Printf("Server is in transition state (%s), cannot toggle.", srvInfo.FormattedStatus())
				return
			}

			ControlsLostFocus() // Para evitar que el botón quede enfocado tras el click
			SetServerInfo(srvInfo.String())
		}(btn)
	})

	log.SetOutput(LogWriter) // Redirigir el log global a nuestro LogWriter
	defer log.SetOutput(nil) // Restablece al valor predeterminado (stderr)

	time.AfterFunc(time.Second, func() {
		LogWriter.Println("✅ [green]TUI En ejecución[-]")
	})
	if err := App.Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
