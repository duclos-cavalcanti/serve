package org

import(
    "fmt"
    "os"
    "log"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-org/cmd/org/config"
)

func display(conf config.Config) {
    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }
    // Set default text style
    defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
    s.SetStyle(defStyle)

    // Clear screen
    s.Clear()
}

func Start(mode string, config_dir string) {
    if mode == "default" {
        conf, err := config.Setup(config_dir)
        if err != nil {
            log.Fatal(err)
        }
        conf.Print()
    } else {
        log.Fatal(fmt.Errorf("Mode: %s is Not defined", mode))
    }
    os.Exit(0)
}
