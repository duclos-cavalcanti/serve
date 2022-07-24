package term

import (
    "log"

	"github.com/gdamore/tcell"
)

type TerminalContext struct {
    Screen tcell.Screen
    DefaultStyle tcell.Style
    Row, Col int
    Filled int
}

func NewStyle(bg, fg tcell.Color) (tcell.Style) {
    return tcell.StyleDefault.Background(bg).Foreground(fg)
}

func NewTerminalContext() (TerminalContext) {
    tc := TerminalContext {
        Row: 0,
        Col: 0,
        Filled: 0,
        DefaultStyle: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite),
        Screen: nil,
    }

    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }

    if err = s.Init(); err != nil {
    		log.Fatalf("%+v", err)
    }

    tc.Screen = s
    tc.Screen.SetStyle(tc.DefaultStyle)
    return tc
}

func (tc *TerminalContext) ClearScreen() {
    tc.Row = 0
    tc.Col = 0
    tc.Filled = 0
    tc.Screen.Clear()
}

func (tc *TerminalContext) Size() (int, int) {
    return tc.Screen.Size()
}

func (tc *TerminalContext) NewLine() {
    tc.Row = 0
    tc.Col++
}

func (tc *TerminalContext) MarkRow() {
    tc.Filled++
}
