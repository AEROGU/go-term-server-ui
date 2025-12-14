package serverinfo

import (
	"fmt"
)

type ServerStatus uint8

const (
	ServerStatusStoped ServerStatus = iota
	ServerStatusStarting
	ServerStatusRunning
	ServerStatusStopping
)

type Info struct {
	IP     string
	Port   int
	Status ServerStatus
}

// String implements the fmt.Stringer interface and returns a
// formatted string representation of the ServerInfo.
// Is formatted with color tags for tview TextView.
// See in tcell.ColorNames the map `ColorNames` for supported color tags. (https://github.com/gdamore/tcell/blob/v2.8.1/color.go)
// Colors use [-] to reset to default color and [colorname] to set a color, e.g., '[red]' for red color.
// See in tview/strings.go the map `attrs` of type map[byte]tcell.AttrMask for supported attribute tags. (https://github.com/rivo/tview/blob/master/strings.go),
// the attributes use lowercase for activationd and uppercase for deactivation, e.g., '[::b]' for bold on and '[::B]' for bold off.
// The format is [foreground:background:attributes] for example [red::b] sets bold and red text withouth changing the background.
func (info Info) String() string {
	return fmt.Sprintf("IP: [yellow]%s[-]\nPort: [yellow]%d[-]\nStatus: %s", info.IP, info.Port, info.FormattedStatus())
}

func (info Info) FormattedStatus() string {
	statusNames := []string{
		ServerStatusStoped:   "[red]Stopped[-]",
		ServerStatusStarting: "[yellow::l]Starting...[-::L]",
		ServerStatusRunning:  "[green]Running[-]",
		ServerStatusStopping: "[orange::l]Stopping...[-::L]",
	}

	status := "[red]UNKNOWN STATUS[-]"
	if int(info.Status) < len(statusNames) {
		status = statusNames[int(info.Status)]
	}

	return status
}
