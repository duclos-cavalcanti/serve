package main

import (
	"github.com/duclos-cavalcanti/go-menu/cmd/menu"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/options"
)

func main() {
    menu.Start(options.ParseOptions())
}
