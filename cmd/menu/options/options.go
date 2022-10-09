package options

import (
    "os"
    "bufio"
    "strings"
)

type Flags struct {
    Mode string
    Prompt string
    Options []string
    delimiter string
}

const usage = `usage: serve [options]
    -m, --mode      mode in which serve will execute, currently only one
    -o, --options   options to select from
    -p, --prompt    menu prompt
    -d, --delimiter string or char used to separate options, default is '\n'
    -h, --help      print usage
`

func ParseOptions() Flags {
    var fs Flags

    fs.Mode = "default"
    fs.delimiter = "\\n"

    args := os.Args[1:]

    for i:=0; i < len(args); i++ {
        arg := args[i]
        switch arg {
        case "-h", "--help":
            os.Stdout.WriteString(usage + "\n")
            os.Exit(1)

        case "-m", "--mode":
            fs.Mode = args[i+1]
            i++

        case "-d", "--delimiter":
            fs.delimiter = args[i+1]
            i++

        case "-p", "--prompt":
            fs.Prompt = args[i+1]
            i++

        case "-o", "--options":
            if i == len(args) - 1 {
                os.Stderr.WriteString("No options provided\n")
                os.Exit(1)
            } else {
                fs.Options = strings.Split(strings.Trim(args[i + 1], "\t "), fs.delimiter)
                i++
            }
            break

        default:
            os.Stderr.WriteString("unknown option: " + arg + "\n")
            os.Exit(1)
        }

    }

    if len(fs.Options) == 0 {
        fi, _ := os.Stdin.Stat()
        if (fi.Mode() & os.ModeCharDevice) == 0 { // serve is used in command chaining, pipe |
            scanner := bufio.NewScanner(os.Stdin)
            i := 0
            for scanner.Scan() {
                fs.Options = append(fs.Options, scanner.Text())
                i++
            }
        } else {
            os.Stderr.WriteString("No options provided\n")
            os.Exit(1)
        }
    }

    if fs.Prompt == "" {
        fs.Prompt = "Menu"
    }

    return fs
}
