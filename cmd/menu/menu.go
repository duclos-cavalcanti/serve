package menu

import(
    "fmt"
    "os"
    "log"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-org/cmd/org/term"
)

func defaultMode(fs Flags) {
    // opts_flag := fs.OptFlag

    tc, err := term.NewTerminalContext()
    if err != nil {
        log.Fatalf("%+v", err)
    }

    if err := tc.Screen.Init(); err != nil {
    		log.Fatalf("%+v", err)
    }

    tc.Screen.SetStyle(tc.DefaultStyle)
    tc.Screen.Clear() // clears screen

    term.DrawTextNewLine(&tc, tc.DefaultStyle, fmt.Sprintf("User: %s", "Example"))
    term.DrawText(&tc, tc.DefaultStyle, fmt.Sprintf("Row %d and Col %d", tc.Row, tc.Col))

    for {
        // Update Screen
        tc.Screen.Show()

        // Poll Event
        ev := tc.Screen.PollEvent()

        // Process event
        switch ev := ev.(type) {
            case *tcell.EventResize:
                // w, h := ev.Size()
                tc.Screen.Sync()

            case *tcell.EventKey:
                if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                tc.Screen.Fini()
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
