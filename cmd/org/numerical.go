package org

import(
    "fmt"
    "strings"
    "strconv"
    "regexp"

	"github.com/duclos-cavalcanti/go-org/cmd/file"
)

const NUMBERS = "0123456789"

func stripNumber(str string) string {
    re := regexp.MustCompile("[0-9]+")
    return re.FindString(str)
}

func containsNumber(str string) bool {
    return strings.ContainsAny(str, NUMBERS)
}

func numerical(dir string) error {
    var err error
    var action_files []file.File
    files, err := file.ReadFiles(dir)

    if err != nil {
        return err
    } else {
        fmt.Println("Mode: Numerical")
        fmt.Println("Directory: ", dir)
        for i,f := range files {
            if containsNumber(f.Name) {
                str := stripNumber(f.Name)
                files[i].Key, err = strconv.Atoi(str)
                if err != nil {
                    return err
                }
                action_files = append(action_files, files[i])
            }
        }

        for _,f := range action_files {
            f.Print()
        }
    }

    return nil
}
