package menu

import (
    "fmt"

	"github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/term"

	"github.com/gdamore/tcell"
)

func parseEvents(app *App, state_channel chan states.State) {
    defer app.wait_group.Done()
    s := app.state
    state_channel <- s

    for {
        // Poll Event
        ev := app.tc.Screen.PollEvent()
        s = app.state

        // Process event
        switch ev := ev.(type) {
            case *tcell.EventResize:
                // w, h := ev.Size()
                app.tc.Screen.Sync()

            case *tcell.EventKey:
                switch ev.Key() {
                    case tcell.KeyEscape, tcell.KeyCtrlC:
                        close(state_channel)
                        return

                    case tcell.KeyRune:
                        switch ev.Rune() {
                            case 'j':
                                if (s.Selected < s.Size) {
                                    s.Selected++
                                }
                            case 'k':
                                if (s.Selected > 0) {
                                    s.Selected--
                                }
                        }
                }
        state_channel <- s
        }
    }
}

func displayApplication(app *App, state_channel <-chan states.State) {
    defer app.wait_group.Done()
    menu_style := term.NewStyle(tcell.ColorBlack, tcell.ColorDefault)
    selected_style := term.NewStyle(tcell.ColorDefault, tcell.ColorRed)
    normal_style := app.tc.DefaultStyle

    for {
        s, more := <- state_channel
        if !more {
            break
        }

        // Clear Screen
        app.tc.ClearScreen()

        // Menu
        term.DrawTextNewLine(app.tc, menu_style, fmt.Sprintf("-- MENU -- [%d]", s.Selected))

        // Options
        for i := range s.Options {
            if i == s.Selected {
                term.DrawTextNewLine(app.tc, selected_style, s.Options[i])
            } else {
                term.DrawTextNewLine(app.tc, normal_style, s.Options[i])
            }
        }

        // Update Screen
        app.tc.Screen.Show()
    }
    return
}
