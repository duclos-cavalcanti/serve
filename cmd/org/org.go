package org

import(
    // "os"
    "fmt"

	"github.com/duclos-cavalcanti/go-org/cmd/file"
)

func numerical(dir string) {
    fmt.Println("numerical")
    var f = file.File {
        Key: 1,
        Name: "foo",
        FileName: "bar",
        PrevFileName: "baz",
    }

    f.Print()
}

func Start(mode string, dir string) {
    if mode == "numerical" {
        numerical(dir)
    } else {
        return
    }
}
