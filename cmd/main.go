package main

import (
    "flag"

	"github.com/duclos-cavalcanti/go-org/cmd/org"
)

func main() {
    var modeFlag = flag.String("mode", "numerical", "mode to perform")
    var dirFlag = flag.String("dir", ".", "target directory")
    // var actionFlag = flag.String("action", "default", "action to perform within the given mode")

    flag.Parse()
    org.Start(*modeFlag, *dirFlag)
}
