package main

import (
	"github.com/duclos-cavalcanti/go-menu/cmd/menu"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/flags"
)

func main() {
    flags := flags.ParseFlags()
    menu.Start(flags)
}
