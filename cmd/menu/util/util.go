package util

import(
    "regexp"
    "strings"
)

const NUMBERS = "0123456789"

func StripNumber(str string) string {
    re := regexp.MustCompile("[0-9]+")
    return re.FindString(str)
}

func ContainsNumber(str string) bool {
    return strings.ContainsAny(str, NUMBERS)
}

func IsEven(val int) bool {
    return val % 2 == 0
}

func IsOdd(val int) bool {
    return val % 2 != 0
}

