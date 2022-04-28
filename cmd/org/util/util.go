package util

import(
    "regexp"
    "strings"
)

const NUMBERS = "0123456789"

func stripNumber(str string) string {
    re := regexp.MustCompile("[0-9]+")
    return re.FindString(str)
}

func containsNumber(str string) bool {
    return strings.ContainsAny(str, NUMBERS)
}

