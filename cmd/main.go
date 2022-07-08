package main

import (
	"github.com/duclos-cavalcanti/go-org/cmd/menu"
)

func main() {
    flags := menu.ParseFlags()
    menu.Start(flags)
}
