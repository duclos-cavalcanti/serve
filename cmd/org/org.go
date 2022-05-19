package org

import(
    "fmt"
    "os"
    "log"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-org/cmd/org/config"
)

func display(conf config.Config) {
    defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite)
    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }

    if err := s.Init(); err != nil {
    		log.Fatalf("%+v", err)
    	}

    s.SetStyle(defStyle)
    s.Clear()

    s.SetContent(0, 0, 'H', nil, defStyle)
    s.SetContent(1, 0, 'i', nil, defStyle)
    s.SetContent(2, 0, '!', nil, defStyle)

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
}

func Start(mode string, config_dir string) {
    if mode == "default" {
        conf, err := config.Setup(config_dir)
        if err != nil {
            log.Fatal(err)
        }
        display(conf)
    } else {
        log.Fatal(fmt.Errorf("Mode: %s is Not defined", mode))
    }
    os.Exit(0)
}
