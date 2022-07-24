package states

import (
    "strings"
)

type State struct {
    Options []string
    Selected int
    Size int
    Chosen bool
}

func CreateState(opts string) State {
    options := strings.SplitN(opts, " ", -1)
    s := State {
        Options: options,
        Size: len(options),
        Selected: 0,
        Chosen: false,
    }

    return s
}
