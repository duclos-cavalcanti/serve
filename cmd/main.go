package main

import (
	"github.com/duclos-cavalcanti/go-org/cmd/org"
)

func main() {
    flags := org.ParseFlags()
    org.Start(flags)
}
