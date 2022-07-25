package term

import (
	"github.com/gdamore/tcell"
)

func (tc *TerminalContext) DrawText(x1, y1 int, style tcell.Style, text string) error {
	row := y1
	col := x1

    for _, c := range []rune(text) {
    	tc.Screen.SetContent(col, row, c, nil, style)
        col++
    }
    return nil
}

func (tc *TerminalContext) DrawBoxWithCorneredText(x1, y1, w, h int, style tcell.Style, text string) {
    s := tc.Screen
    x2 := x1 + w - 1
    y2 := y1 + h - 1
    text_x := x1 + 1
    text_y := y1 + 1

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw Horizontal borders
	for col := x1 + 1; col < x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}

	// Draw Vertical borders
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Draw Corners
	s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)

    tc.DrawText(text_x, text_y, style, text)
}

func (tc *TerminalContext) DrawBox(x1, y1, w, h int, style tcell.Style) {
    s := tc.Screen
    x2 := x1 + w - 1
    y2 := y1 + h - 1

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw Horizontal borders
	for col := x1 + 1; col < x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}

	// Draw Vertical borders
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Draw Corners
	s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
}
