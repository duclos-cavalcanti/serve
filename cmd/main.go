package main

import (
    "flag"

	"github.com/duclos-cavalcanti/go-org/cmd/org"
)

func main() {
    var modeFlag = flag.String("mode", "numerical", "mode to perform")
    var dirFlag = flag.String("dir", ".", "target directory")

    flag.Parse()

    org.Start(*modeFlag, *dirFlag)
}
