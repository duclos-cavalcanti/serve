package main

import (
    "flag"

	"github.com/duclos-cavalcanti/go-org/cmd/org"
)

func main() {
    var modeFlag = flag.String("mode", "default", "mode to perform")
    var configFlag = flag.String("config", "~/.config/go-org", "config directory")

    flag.Parse()
    org.Start(*modeFlag, *configFlag)
}
