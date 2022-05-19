package util

import(
    "os"
    "io/ioutil"
)

func ReadFile(file string) ([]byte, error) {
    out,err := ioutil.ReadFile(file)
    return out,err
}

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

func IsDirectory(dir string) bool {
    file_info, err := os.Stat(dir)
    if err != nil {
        return false
    } else {
        if file_info.IsDir() {
            return true
        } else {
            return false
        }
    }
}
