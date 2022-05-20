package util

import(
    "regexp"
    "strings"
    "unicode/utf8"
)

const NUMBERS = "0123456789"

func LengthString(s string) int {
    var cnt = utf8.RuneCountInString(s)
    return cnt
}

func StripNumber(str string) string {
    re := regexp.MustCompile("[0-9]+")
    return re.FindString(str)
}

func ContainsNumber(str string) bool {
    return strings.ContainsAny(str, NUMBERS)
}

