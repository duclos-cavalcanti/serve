package file

import (
    "os"
    "io/ioutil"
    "fmt"
    "strings"
    // "errors"
)

type File struct {
    Key int
    Name, FileType string
}

func (f File) Print() {
    fmt.Println("Name: ", f.Name)
    fmt.Println("Type: ", f.FileType)
    fmt.Println("Key: ", f.Key)
    fmt.Printf("\n")
}

func ReadFiles(dir string) ([]File, error) {
    var files []File
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
            var f_ = File {
                Key: -1,
                Name: name,
                FileType: end,
            }
            files = append(files, f_)
        }
    }

    return files, nil
}

