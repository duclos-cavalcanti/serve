package util

import(
    "fmt"
    "os"
    "io/ioutil"
    "errors"
)

func WriteFile(name string, in []byte) error {
    err := ioutil.WriteFile(name, in, 0644)
    if err != nil {
        return err
    } else {
        return nil
    }
}

func ExistsFile(file string) bool {
    _, err := os.Open(file)
    if err != nil {
        return false
    } else {
        return true
    }
}

func IsDirectory(dir string) (bool, error) {
    file_info, err := os.Stat(dir)
    if err != nil {
        return false, err
    } else {
        if file_info.IsDir() {
            return true, nil
        } else {
            return false, errors.New(fmt.Sprintf("%s is not a valid directory", dir))
        }
    }
}
