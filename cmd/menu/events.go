package menu

import (
    "fmt"

	"github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/term"

	"github.com/gdamore/tcell"
)

func parseEvents(app *App, state_channel chan states.State, debug_channel chan string) {
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
                        close(debug_channel)
                        return

                    case tcell.KeyRune:
                        switch ev.Rune() {
                            case 'j':
                                debug_channel <- fmt.Sprintf("J has been pressed")
                                if (s.Selected < s.Size) {
                                    s.Selected++
                                    debug_channel <- fmt.Sprintf("Sel incremented: %d", s.Selected)
                                }
                            case 'k':
                                debug_channel <- fmt.Sprintf("K has been pressed")
                                if (s.Selected > 0) {
                                    s.Selected--
                                    debug_channel <- fmt.Sprintf("Sel decremented: %d", s.Selected)
                                }
                        }
                }
        state_channel <- s
        }
    }
}

func displayApplication(app *App, state_channel <-chan states.State, debug_channel chan string) {
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
        term.DrawTextNewLine(app.tc, menu_style, fmt.Sprintf("-- MENU --"))

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
