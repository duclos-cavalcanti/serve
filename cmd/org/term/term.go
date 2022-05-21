package term

import (
	// "github.com/gdamore/tcell"
)

type TerminalContext struct {
    row, col, width, height int
}

func NewTerminalContext() TerminalContext {
    var s TerminalContext
    s.row = 0
    s.col = 0
    s.width = 0
    s.height = 0

    return s
}

func (s TerminalContext) Row() int {
    return s.row
}

func (s TerminalContext) Col() int {
    return s.col
}

func (s TerminalContext) Size() (int, int) {
    return s.width, s.height
}

func (s TerminalContext) SetWidth(w int) {
    s.width = w
}

func (s TerminalContext) SetHeight(h int) {
    s.height = h
}
