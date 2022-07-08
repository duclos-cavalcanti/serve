package menu

import (
    "fmt"
    "flag"
    "os"
)

type Flags struct {
    ModeFlag, OptFlag string
}


func ParseFlags() Flags {
    var fs Flags
    var modeFlag = flag.String("mode", "default", "mode to perform")
    var optFlag = flag.String("options", "", "options to select from")

    flag.Parse()

    fs.ModeFlag = *modeFlag
    fs.OptFlag = *optFlag

    fmt.Println("%s", *optFlag)

    os.Exit(0)
    return fs
}
