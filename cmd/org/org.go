package org

import(
    "fmt"
    "log"
    "errors"

    "github.com/duclos-cavalcanti/go-org/cmd/org/config"
)

func Start(mode string, conf string) {
    var err error
    if mode == "default" {
        err = config.Setup(conf)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        err = errors.New(fmt.Sprintf("Mode: %s is Not defined", mode))
        log.Fatal(err)
    }
}
