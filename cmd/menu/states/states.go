package states

import (
	"github.com/duclos-cavalcanti/go-menu/cmd/menu/options"
)

type State struct {
    Options []string
    Prompt string
    Selected int
    Size int
    Chosen bool
}

func CreateState(fs options.Flags) State {
    var prompt string
    options := fs.Options
    prompt = fs.Prompt

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
