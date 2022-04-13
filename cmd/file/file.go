package file

import (
    "fmt"
)

type File struct {
    Key int
    Name, FileName, PrevFileName string
}

func (f File) Print() {
    fmt.Printf("%+v", f)
}
