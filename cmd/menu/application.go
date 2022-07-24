package menu

import (
    "sync"

    "github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
    "github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

type App struct {
    tc *term.TerminalContext
    state *states.State
    wait_group *sync.WaitGroup
}

func CreateApp(tc *term.TerminalContext,
               wg *sync.WaitGroup) App {
    a := App {
        tc: tc,
        wait_group: wg,
    }

    return a
}
