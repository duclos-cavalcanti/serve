package flags

import (
    // "fmt"
    // "os"
    "flag"
)

type Flags struct {
    ModeFlag, OptFlag, PromptFlag string
}


func ParseFlags() Flags {
    var fs Flags
    var modeFlag = flag.String("mode", "default", "mode to perform")
    var optFlag = flag.String("options", "", "options to select from")
    var promptFlag = flag.String("prompt", "Menu", "menu prompt")

    flag.Parse()

    fs.ModeFlag = *modeFlag
    fs.OptFlag = *optFlag
    fs.PromptFlag = *promptFlag

    return fs
}
