package file

import (
    "os"
    "io/ioutil"
    "fmt"
    // "errors"
)

type File struct {
    Key int
    Name, FileName, PrevFileName string
}

func (f File) Print() {
    fmt.Printf("%+v", f)
}

func ReadFiles(dir string) ([]File, error) {
    var files []File
    var err error

    // change directory if needed
    if dir != "." {
        err := os.Chdir(dir)
        if err != nil {
            return files, err
        }
    }

    // read current directory value
    cwd, err := os.Getwd()
    if err != nil {
        return files, err
    }

    // obtain listed files
    fs, err := ioutil.ReadDir(cwd)
    for i, f := range fs {
        var f_ = File {
            Key: i,
            Name: "",
            FileName: f.Name(),
            PrevFileName: f.Name(),
        }

        files = append(files, f_)
    }

    return files, nil
}

