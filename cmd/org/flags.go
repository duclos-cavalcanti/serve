package org

import (
    "flag"
)

type Flags struct {
    ModeFlag, ConfigFlag string
}


func ParseFlags() Flags {
    var fs Flags
    var modeFlag = flag.String("mode", "default", "mode to perform")
    var configFlag = flag.String("config", "~/.config/go-org", "config directory")

    flag.Parse()

    fs.ModeFlag = *modeFlag
    fs.ConfigFlag = *configFlag

    return fs
}
