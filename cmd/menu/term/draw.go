package term

import (
    "errors"

	"github.com/gdamore/tcell"
    "github.com/duclos-cavalcanti/go-menu/cmd/menu/util"
)

func DrawText(tc *TerminalContext, style tcell.Style, text string) error {
	row := tc.Row
    w, _ := tc.Size()
    length := util.LengthString(text)

    if length + row > w {
        return errors.New("Text length would exceed width.")
    } else {
	    for _, c := range []rune(text) {
	    	tc.Screen.SetContent(tc.Row, tc.Col, c, nil, style)
            tc.Row++
	    }
        tc.MarkRow()
        return nil
    }
}

func DrawTextNewLine(tc *TerminalContext, style tcell.Style, text string) error {
    _, h := tc.Size()
    if tc.Col + 1 >= h {
        return errors.New("NewLine would exceed available columns in terminal")
    } else {
        err := DrawText(tc, style, text)
        if err != nil {
            return err
        } else {
            tc.NewLine()
            return nil
        }
    }
}

func OffsetDrawnText(tc *TerminalContext, x0, x1, y, offset int) error {
    diff := x1 - x0 + 1
    str := make([]rune, diff)
    styles := make([]tcell.Style, diff)
    w, h := tc.Size()

    if y >= h {
        return errors.New("Column exceeds available height")
    } else if x1 >= w {
        return errors.New("Row exceeds available width")
    } else {
        for i := 0; i<=diff; i++ {
            str[i], _, styles[i], _ = tc.Screen.GetContent(x0 + i, y)
            tc.Screen.SetContent(x0 + i, y, ' ', nil, tc.DefaultStyle)
        }

        for i := 0; i<=diff; i++ {
            tc.Screen.SetContent(x0 + i + 2, y, str[i], nil, styles[i])
        }

        return nil
    }
}
