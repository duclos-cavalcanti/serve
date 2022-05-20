package org

import(
    "fmt"
    "os"
    "log"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-org/cmd/org/term"
    "github.com/duclos-cavalcanti/go-org/cmd/org/config"
)

func Start(mode string, config_dir string) {
    if mode == "default" {
        _, err := config.Setup(config_dir)
        if err != nil {
            log.Fatal(err)
        }

        defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite)

        sc := term.NewScreenContext()
        s, err := tcell.NewScreen()
        if err != nil {
            log.Fatalf("%+v", err)
        }

        if err := s.Init(); err != nil {
        		log.Fatalf("%+v", err)
        }

        s.SetStyle(defStyle)
        s.Clear()

        term.DrawText(s, sc, defStyle, "Hi!")

        for {
            // Update Screen
            s.Show()

            // Poll Event
            ev := s.PollEvent()

            // Process event
            switch ev := ev.(type) {
                case *tcell.EventResize:
                    s.Sync()

                case *tcell.EventKey:
                    if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                    s.Fini()
                    os.Exit(0)
                }
            }
        }
    } else {
        log.Fatal(fmt.Errorf("Mode: %s is Not defined", mode))
    }
    os.Exit(0)
}
