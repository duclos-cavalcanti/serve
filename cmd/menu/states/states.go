package states

import (
    "strings"
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/flags"
)

type State struct {
    Options []string
    Prompt string
    Selected int
    Size int
    Chosen bool
}

func CreateState(fs flags.Flags) State {
    var prompt string
    options := strings.Split(strings.Trim(fs.OptFlag, "\t "), `\n`)
    prompt = fs.PromptFlag

    s := State {
        Options: options,
        Prompt: prompt,
        Size: len(options),
        Selected: 0,
        Chosen: false,
    }

    return s
}

func (s *State) MaxOptionSize() int {
    max := 0
    for i := range s.Options {
        val := len(s.Options[i])
        if (val > max) {
            max = val
        }
    }

    return max
}

func (s *State) PromptSize() int {
    return len(s.Prompt)
}
