package menu

import (
    "sync"

    "github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
    "github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

type App struct {
    tc *term.TerminalContext
    state states.State
    wait_group *sync.WaitGroup
}

func CreateApp(tc *term.TerminalContext, initial_state states.State, wg *sync.WaitGroup) App {
    a := App {
        tc: tc,
        state: initial_state,
        wait_group: wg,
    }

    return a
}

func (app *App) hasUserChosen() bool {
    return app.state.Chosen
}
