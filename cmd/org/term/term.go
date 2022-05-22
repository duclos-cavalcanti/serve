package term

import (
	"github.com/gdamore/tcell"
)

type TerminalContext struct {
    Screen tcell.Screen
    DefaultStyle tcell.Style
    Row, Col int
    width, height int
}

func NewTerminalContext() (TerminalContext, error) {
    tc := TerminalContext {
        Row: 0,
        Col: 0,
        width: 0,
        height: 0,
        DefaultStyle: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite),
        Screen: nil,
    }

    s, err := tcell.NewScreen()
    w, h := s.Size()
    tc.width = w
    tc.height = h
    tc.Screen = s

    return tc, err
}

func (tc *TerminalContext) Size() (int, int) {
    return tc.Screen.Size()
}

func (tc *TerminalContext) Width() int {
    return tc.width
}

func (tc *TerminalContext) Height() int {
    return tc.height
}

func (tc *TerminalContext) SetSize(w, h int) {
    tc.width = w
    tc.height = h
}

func (tc *TerminalContext) NewLine() {
    tc.Row = 0
    tc.Col++
}
