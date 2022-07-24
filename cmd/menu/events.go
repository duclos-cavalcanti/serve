package menu

import (
	"github.com/gdamore/tcell"

	"github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

func parseEvents(app *App, state_channel chan states.State) {
    defer app.wait_group.Done()
    for {
        // Poll Event
        ev := app.tc.Screen.PollEvent()
        s := app.state

        // Process event
        switch ev := ev.(type) {
            case *tcell.EventResize:
                // w, h := ev.Size()
                app.tc.Screen.Sync()

            case *tcell.EventKey:
                if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                    close(state_channel)
                    break
            }
        }
        state_channel <- *s
    }
    return
}

func displayApplication(app *App, state_channel <-chan states.State) {
    defer app.wait_group.Done()
    for {
        s, more := <- state_channel
        if !more {
            break
        }

        for i := range s.Options {
            term.DrawTextNewLine(app.tc, app.tc.DefaultStyle, s.Options[i])
        }

        // Update Screen
        app.tc.Screen.Show()
    }
    return
}
