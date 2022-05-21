package org

import(
    "fmt"
    "os"
    "log"

	"github.com/gdamore/tcell"

    "github.com/duclos-cavalcanti/go-org/cmd/org/term"
    "github.com/duclos-cavalcanti/go-org/cmd/org/config"
)

func defaultMode(fs Flags) {
    config_flag := fs.ConfigFlag
    _, err := config.Setup(config_flag)
    if err != nil {
        log.Fatal(err)
    }

    tc := term.NewTerminalContext()
    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }

    if err := s.Init(); err != nil {
    		log.Fatalf("%+v", err)
    }

    defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite)
    s.SetStyle(defStyle)
    s.Clear() // clears screen

    term.DrawText(s, tc, defStyle, "Hi!")

    for {
        // Update Screen
        s.Show()

        // Poll Event
        ev := s.PollEvent()

        // Process event
        switch ev := ev.(type) {
            case *tcell.EventResize:
                w, h := ev.Size()
                tc.SetWidth(w)
                tc.SetHeight(h)
                s.Sync()

            case *tcell.EventKey:
                if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                s.Fini()
                os.Exit(0)
            }
        }
    }
}

func Start(fs Flags) {
    if fs.ModeFlag == "default" {
        defaultMode(fs)
    } else {
        log.Fatal(fmt.Errorf("Mode: %s is Not defined", fs.ModeFlag))
    }
    os.Exit(0)
}
