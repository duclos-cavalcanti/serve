package states

import (
    "strings"
)

type State struct {
    Options []string
    Selected int
}

func NewState(opts string) State {
    s := State {
        Options: strings.SplitN(opts, " ", -1),
        Selected: 0,
    }

    return s
}
