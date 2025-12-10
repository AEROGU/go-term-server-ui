package main

import (
	"log"
	"time"
)

func main() {
	InitializeUI()

	// Test: Actualizar info del servidor después de 5 segundos -----------------
	time.AfterFunc(5*time.Second, func() {
		SetServerInfo("Test")
	})

	if err := App.Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
