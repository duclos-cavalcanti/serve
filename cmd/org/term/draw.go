package term

import (
    "errors"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-org/cmd/org/util"
)

func DrawText(s tcell.Screen, sc ScreenContext, style tcell.Style, text string) error {
	row := sc.Row()
    w, _ := s.Size()
    length := util.LengthString(text)

    if length + row > w {
        return errors.New("Unable to draw to screen, text length would exceed width.")
    } else {
	    for _, c := range []rune(text) {
	    	s.SetContent(sc.row, sc.col, c, nil, style)
            sc.row++
	    }
        return nil
    }
}
