package menu

import (
	"log"
	"os"
	"sync"

	"github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

var (
    debug_channel chan string
    state_channel chan states.State
)


func defaultMode(fs Flags) {
    var wait_group sync.WaitGroup

    tc := term.NewTerminalContext()

    debug_channel = make(chan string)
    state_channel = make(chan states.State)

    application := CreateApp(&tc,
                             states.CreateState(fs.OptFlag),
                             &wait_group)

    wait_group.Add(3)
    go parseEvents(&application)
    go displayApplication(&application)
    go logger(application.wait_group)
    wait_group.Wait()

    tc.Screen.Fini()

    if (application.hasUserChosen()) {
        println(application.state.Options[application.state.Selected])
        os.Exit(0)
    } else {
        os.Exit(-1)
    }
}

func Start(fs Flags) {
    if fs.ModeFlag == "default" {
        defaultMode(fs)
    } else {
        log.Fatalf("Mode: %s is Not defined", fs.ModeFlag)
    }
    os.Exit(0)
}
