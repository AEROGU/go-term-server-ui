package main

import "github.com/rivo/tview"

// ControlsAddButton agrega un nuevo botón a la sección de controles con la etiqueta y función proporcionadas.
// Devuelve una referencia al botón creado para posibles manipulaciones futuras.
// Recuerda llamar a RefreshUI() si deseas que los cambios se reflejen inmediatamente en la UI.
// También puedes usar RefreshUI() después de agregar múltiples botones para optimizar el rendimiento,
// o llamar a App.QueueUpdateDraw() directamente si estás dentro de una gorutina.
func ControlsAddButton(label string, selectedFunc func()) *tview.Button {
	button := tview.NewButton(label)
	if selectedFunc != nil {
		button.SetSelectedFunc(selectedFunc)
	}
	controlsFlex.AddItem(button, 0, 1, false)
	return button
}

// ControlsAddItem agrega un item genérico (puede ser un botón u otro primitive) a la sección de controles.
func ControlsAddItem(item tview.Primitive, fixedSize int, proportion int, focus bool) {
	controlsFlex.AddItem(item, fixedSize, proportion, focus)
}

// ControlsRemoveItem elimina un botón o item específico de la sección de controles.
func ControlsRemoveItem(item tview.Primitive) {
	if item != nil {
		controlsFlex.RemoveItem(item)
	}
}

func ControlsLostFocus() {
	if App != nil {
		App.SetFocus(nil)
	}
}
