package term

import (
	"github.com/gdamore/tcell"
)

type TerminalContext struct {
    Screen tcell.Screen
    DefaultStyle tcell.Style
    Row, Col int
}

func NewTerminalContext() (TerminalContext, error) {
    tc := TerminalContext {
        Row: 0,
        Col: 0,
        DefaultStyle: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite),
        Screen: nil,
    }

    s, err := tcell.NewScreen()
    tc.Screen = s

    return tc, err
}

func (tc *TerminalContext) Size() (int, int) {
    return tc.Screen.Size()
}

func (tc *TerminalContext) NewLine() {
    tc.Row = 0
    tc.Col++
}
