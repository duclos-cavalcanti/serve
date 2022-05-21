package term

import (
    "errors"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-org/cmd/org/util"
)

func DrawText(s tcell.Screen, tc TerminalContext, style tcell.Style, text string) error {
	row := tc.Row()
    w, _ := s.Size()
    length := util.LengthString(text)

    if length + row > w {
        return errors.New("Unable to draw to screen, text length would exceed width.")
    } else {
	    for _, c := range []rune(text) {
	    	s.SetContent(tc.row, tc.col, c, nil, style)
            tc.row++
	    }
        return nil
    }
}
