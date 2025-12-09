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
