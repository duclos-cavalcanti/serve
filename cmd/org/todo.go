package org

import (
    "os"
    "io/ioutil"
    "fmt"
    "strings"
)

type notebook struct {
    Key int
    Name, FileType string
}

func (f notebook) Print() {
    fmt.Println("Name: ", f.Name)
    fmt.Println("Type: ", f.FileType)
    fmt.Println("Key: ", f.Key)
    fmt.Printf("\n")
}

func ReadFiles(dir string) ([]notebook, error) {
    var files []notebook
    var err error

    // change directory
    err = os.Chdir(dir)
    if err != nil {
        return files, err
    }

    // read current directory
    cwd, err := os.Getwd()
    if err != nil {
        return files, err
    }

    // read files
    fs, err := ioutil.ReadDir(cwd)
    for _, f := range fs {
        var info, err = os.Stat(f.Name())
        if ! os.IsNotExist(err) && !info.IsDir() {
            var name, end, _ = strings.Cut(f.Name(), ".")
            var f_ = notebook {
                Key: -1,
                Name: name,
                FileType: end,
            }
            files = append(files, f_)
        }
    }

    return files, nil
}

