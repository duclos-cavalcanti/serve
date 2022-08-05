package menu

import (
	"log"
	"os"
	"sync"
    "fmt"

	"github.com/duclos-cavalcanti/go-menu/cmd/menu/flags"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

var (
    debug_channel chan string
    state_channel chan states.State
    wait_group sync.WaitGroup
)


func defaultMode(fs flags.Flags) {
    tc := term.NewTerminalContext()

    debug_channel = make(chan string)
    state_channel = make(chan states.State)

    application := CreateApp(&tc, states.CreateState(fs))

    wait_group.Add(3)
    go parseEvents(&application)
    go display(&application)
    go logger()
    wait_group.Wait()

    tc.Screen.Fini()

    if (application.state.Chosen) {
        fmt.Println(application.state.Options[application.state.Selected])
        os.Exit(0)
    } else {
        os.Exit(-1)
    }
}

func Start(fs flags.Flags) {
    if fs.ModeFlag == "default" {
        defaultMode(fs)
    } else {
        log.Fatalf("Mode: %s is Not defined", fs.ModeFlag)
    }
    os.Exit(0)
}
