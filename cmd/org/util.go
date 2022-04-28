package org

import(
    "fmt"
    "os"
    "io/ioutil"
    "regexp"
    "strings"
    "errors"
)

const NUMBERS = "0123456789"

func stripNumber(str string) string {
    re := regexp.MustCompile("[0-9]+")
    return re.FindString(str)
}

func containsNumber(str string) bool {
    return strings.ContainsAny(str, NUMBERS)
}

func writeFile(name string, in []byte) error {
    err := ioutil.WriteFile(name, in, 0644)
    if err != nil {
        return err
    } else {
        return nil
    }
}

func existsFile(file string) bool {
    _, err := os.Open(file)
    if err != nil {
        return false
    } else {
        return true
    }
}

func isDirectory(dir string) (bool, error) {
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
