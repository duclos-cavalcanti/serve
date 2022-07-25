package term

import (
    "log"

	"github.com/gdamore/tcell"
)

type TerminalContext struct {
    Screen tcell.Screen
    DefaultStyle tcell.Style
}

func NewStyle(bg, fg tcell.Color) (tcell.Style) {
    return tcell.StyleDefault.Background(bg).Foreground(fg)
}

func NewTerminalContext() (TerminalContext) {
    tc := TerminalContext {
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
    tc.Screen.Clear()
}

func (tc *TerminalContext) Size() (int, int) {
    return tc.Screen.Size()
}

func (tc *TerminalContext) Center() (int, int) {
    var x_center int
    var y_center int

    w, h := tc.Size()
    x_center = w / 2
    y_center = h / 2

    return x_center, y_center
}
