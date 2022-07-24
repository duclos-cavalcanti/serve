package main

import (
	"github.com/duclos-cavalcanti/go-menu/cmd/menu"
)

func main() {
    flags := menu.ParseFlags()
    menu.Start(flags)
}
