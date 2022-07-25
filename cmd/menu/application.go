package menu

import (
    "sync"
    "fmt"

	"github.com/gdamore/tcell"

    "github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
    "github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

type App struct {
    tc *term.TerminalContext
    state states.State
}

func CreateApp(tc *term.TerminalContext, initial_state states.State, wg *sync.WaitGroup) App {
    a := App {
        tc: tc,
        state: initial_state,
    }

    return a
}

func (app *App) hasUserChosen() bool {
    return app.state.Chosen
}

func parseApplicationEvents(app *App) {
    defer wait_group.Done()
    s := app.state
    state_channel <- s
    logEvent(fmt.Sprintf("Size of options: %d", s.Size))

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

                    case tcell.KeyEnter:
                        app.state.Chosen = true
                        close(state_channel)
                        close(debug_channel)
                        return

                    case tcell.KeyRune:
                        switch ev.Rune() {
                            case 'j':
                                logEvent(fmt.Sprintf("J has been pressed, sel: %d", s.Selected))
                                if (s.Selected < s.Size) {
                                    s.Selected++
                                    logEvent(fmt.Sprintf("Sel incremented: %d", s.Selected))
                                }
                            case 'k':
                                logEvent(fmt.Sprintf("K has been pressed, sel: %d", s.Selected))
                                if (s.Selected > 0) {
                                    s.Selected--
                                    logEvent(fmt.Sprintf("Sel decremented: %d", s.Selected))
                                }
                        }
                }
        state_channel <- s
        app.state = s
        }
    }
}

func displayApplication(app *App) {
    defer wait_group.Done()
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
