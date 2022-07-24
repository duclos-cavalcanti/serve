package states

import (
    "strings"

    "github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

type State struct {
    options []string
    selected int
}

func NewState(opts string) State {
    s := State {
        options: strings.SplitN(opts, " ", -1),
        selected: 0,
    }

    return s
}
