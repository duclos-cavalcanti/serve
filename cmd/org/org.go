package org

import(
    "fmt"

	// "github.com/duclos-cavalcanti/go-org/cmd/file"
)

func numerical() {
    println("numerical")
}

func Start(mode String) {
    if mode == "numerical" {
        numerical()
    } else {
        return
    }
}
