package term

import (
	// "github.com/gdamore/tcell"
)

type ScreenContext struct {
    row, col int
}

func NewScreenContext() ScreenContext {
    var s ScreenContext
    s.row = 0
    s.col = 0

    return s
}

func (s ScreenContext) Row() int {
    return s.row
}

func (s ScreenContext) Col() int {
    return s.col
}
