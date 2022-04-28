package org

import(
    "fmt"
    "log"
    "errors"
)

func Start(mode string, config string) {
    var err error
    if mode == "default" {
        err = setup(config)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        err = errors.New(fmt.Sprintf("Mode: %s is Not defined", mode))
        log.Fatal(err)
    }
}
