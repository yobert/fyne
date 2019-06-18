package gl
/*
import (
	"fmt"
)

// serialEvent encapsulates various event types that need to be dispatched
// serially, to ensure consistency between handlers. For example: Maybe you want a mouse
// click handler that does something only if the shift key is held down. If you press shift and
// very shortly after click the mouse, it is important that the shift key handler has finished
// before the mouse click handler begins.
//
// Another example is for people who type very fast: The key runes need to arrive in the exact
// order they were dispatched. Or people who click fast: The MouseDown handler must finish before
// the MouseUp handler.
//
// So far, supported serialEvent types are:
// *fyne.MouseEvent
// *fyne.KeyEvent
// rune (TypedKey)
*/
type serialEvent struct {
	target interface{}
	event interface{}
}


// each window will allocate a serialEvent channel and send events to it. When the window is
// closed, the channel will be closed and this function will return.
func handleSerialEvents(events <-chan serialEvent) {
}

/*
	for sev := range events {
		switch ev := sev.event.(type) {
		case *fyne.MouseEvent:
			t, ok := sev.target.(desktop.Mouseable)
			if !ok {
				fyne.LogError(fmt.Sprintf("Mouse click event targeted at something not Mouseable: %T", sev.target))
				continue
			}
			
		case *fyne.KeyEvent:
		case rune:
		default:
			fyne.LogError(fmt.Sprintf("Unhandled serialEvent type %T", sev))
		}
	}
}*/
