package menu

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/duclos-cavalcanti/go-menu/cmd/menu/states"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/term"
)

func defaultMode(fs Flags) {
    var wait_group sync.WaitGroup

    initial_state := states.NewState(fs.OptFlag)
    state_channel := make(chan states.State)
    tc := term.NewTerminalContext()

    application := CreateApp(&tc, &wait_group)
    application.state = (*states.State)(&initial_state)

    wait_group.Add(2)
    go parseEvents(&application, state_channel)
    go displayApplication(&application, state_channel)
    wait_group.Wait()

    tc.Screen.Fini()
    os.Exit(0)
}

func Start(fs Flags) {
    if fs.ModeFlag == "default" {
        defaultMode(fs)
    } else {
        log.Fatal(fmt.Errorf("Mode: %s is Not defined", fs.ModeFlag))
    }
    os.Exit(0)
}
