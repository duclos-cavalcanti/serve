package main

import (
    // "os"
	"fmt"
    "flag"

	"github.com/duclos-cavalcanti/go-org/cmd/file"
)

func main() {
    var modeFlag = flag.String("mode", "numerical", "organization mode")

    flag.Parse()

    file.HelloWorld()
    fmt.Println(*modeFlag)
}
