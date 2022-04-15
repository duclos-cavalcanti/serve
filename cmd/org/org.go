package org

import(
    "fmt"
    "log"
    "errors"
)

func Start(mode string, dir string) {
    var err error
    if mode == "numerical" {
        err = numerical(dir)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        err = errors.New(fmt.Sprintf("Mode: %s is Not defined", mode))
        log.Fatal(err)
    }
}
