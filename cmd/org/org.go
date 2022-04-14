package org

import(
    "fmt"

	"github.com/duclos-cavalcanti/go-org/cmd/file"
)

func numerical(dir string) {
    files, err := file.ReadFiles(dir)
    fmt.Println(files)
    if err != nil {
        fmt.Println("Error=: ", err)
    } else {

    }
}

func Start(mode string, dir string) {
    if mode == "numerical" {
        numerical(dir)
    } else {
        return
    }
}
