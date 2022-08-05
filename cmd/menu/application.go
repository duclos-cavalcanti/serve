package menu

import (
    "fmt"

	"github.com/gdamore/tcell"

    "github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
    "github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

const (
    RECTANGLE_SINGLE_HEIGHT = 3
)

type App struct {
    tc *term.TerminalContext
    state states.State
}

func CreateApp(tc *term.TerminalContext, initial_state states.State) App {
    a := App {
        tc: tc,
        state: initial_state,
    }

    return a
}

func parseEvents(app *App) {
    defer wait_group.Done()
    closeChannels := func(){ close(state_channel); close(debug_channel) }

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
                        logEvent("Leaving without chosen option")
                        closeChannels()
                        return

                    case tcell.KeyEnter:
                        app.state.Chosen = true
                        logEvent(fmt.Sprintf("Option chosen: %s", s.Options[s.Selected]))
                        closeChannels()
                        return

                    case tcell.KeyRune:
                        switch ev.Rune() {
                            case 'j':
                                if (s.Selected < s.Size - 1) {
                                    s.Selected++
                                }
                                logEvent(fmt.Sprintf("J has been pressed, sel: %d", s.Selected))
                            case 'k':
                                if (s.Selected > 0) {
                                    s.Selected--
                                }
                                logEvent(fmt.Sprintf("K has been pressed, sel: %d", s.Selected))
                        }
                }
        app.state = s
        state_channel <- s
        }
    }
}

func display(app *App) {
    defer wait_group.Done()
    normal_style := app.tc.DefaultStyle
    selected_style := term.NewStyle(tcell.ColorDefault, tcell.ColorRed)
    menu := app.state.Prompt
    menu_length := app.state.PromptSize()
    max_option_length := app.state.MaxOptionSize()
    for {
        s, more := <- state_channel
        if !more {
            return
        }
        // Clear Screen
        app.tc.ClearScreen()

        center_x, center_y := app.tc.Center()

        menu_width := menu_length + 2
        menu_height := RECTANGLE_SINGLE_HEIGHT

        options_height := app.state.Size + 2
        options_width := max_option_length * 5 + 2 // add 2 due to the offsetted '> '

        menu_corner_x := center_x - menu_width/2
        menu_corner_y := center_y - options_height/2 - menu_height

        options_corner_x := center_x - options_width/2
        options_corner_y := center_y - options_height/2

        // Menu
        app.tc.DrawBoxWithCorneredText(menu_corner_x,
                                       menu_corner_y,
                                       menu_width,
                                       menu_height,
                                       normal_style, menu)

        // Options Box
        app.tc.DrawBox(options_corner_x,
                       options_corner_y,
                       options_width,
                       options_height,
                       normal_style)

        // Options
        options_x := options_corner_x + 1
        options_y := options_corner_y + 1
        for i := range s.Options {
            if i == s.Selected {
                app.tc.DrawText(options_x, options_y, selected_style, fmt.Sprintf("> %s", s.Options[i]))
            } else {
                app.tc.DrawText(options_x, options_y, normal_style, fmt.Sprintf("  %s", s.Options[i]))
            }
            options_y++
        }

        // Update Screen
        app.tc.Screen.Show()
    }
}
