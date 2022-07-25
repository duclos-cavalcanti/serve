package menu

import (
	"log"
	"os"
	"sync"

	"github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/util"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

func defaultMode(fs Flags) {
    var wait_group sync.WaitGroup

    debug_channel := make(chan string)
    state_channel := make(chan states.State)
    tc := term.NewTerminalContext()

    application := CreateApp(&tc,
                             states.CreateState(fs.OptFlag),
                             &wait_group)

    wait_group.Add(3)
    go parseEvents(&application, state_channel, debug_channel)
    go displayApplication(&application, state_channel, debug_channel)
    go util.LogEvents(debug_channel, application.wait_group)
    wait_group.Wait()

    tc.Screen.Fini()

    if (application.hasUserChosen()) {
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
