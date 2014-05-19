package box

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func CoverPrintStr(x, y int, str string, fg, bg termbox.Attribute) error {

	//str := "Win,use ESC exit"
	xx := x
	for n, c := range str {
		if c == '\n' {
			y++
			xx = x - n - 1
		}
		termbox.SetCell(xx+n, y, c, fg, bg)
	}
	termbox.Flush()
	return nil
}

func (t Box) Print(ox, oy int) error {
	fg := termbox.ColorYellow
	bg := termbox.ColorBlack
	termbox.Clear(fg, bg)
	str := []rune("  成 绩  : " + fmt.Sprint(Score))
	for n, c := range str {
		termbox.SetCell(ox+n, oy-1, c, fg, bg)
	}
	str = []rune("ESC:退 出  " + " Enter:重 玩")
	for n, c := range str {
		termbox.SetCell(ox+n, oy-2, c, fg, bg)
	}
	str = []rune("用 方 向 键 ← ↑ ↓ → 游 戏 ! ")
	for n, c := range str {
		termbox.SetCell(ox+n, oy-3, c, fg, bg)
	}
	fg = termbox.ColorBlack
	bg = termbox.ColorGreen
	for i := 0; i <= len(t); i++ {
		for x := 0; x < 5*len(t); x++ {
			termbox.SetCell(ox+x, oy+i*2, '-', fg, bg)
		}
		for x := 0; x <= 2*len(t); x++ {
			if x%2 == 0 {
				termbox.SetCell(ox+i*5, oy+x, '+', fg, bg)
			} else {
				termbox.SetCell(ox+i*5, oy+x, '|', fg, bg)
			}
		}
	}
	fg = termbox.ColorYellow
	bg = termbox.ColorBlack
	for i := range t {
		for j := range t[i] {
			if t[i][j] > 0 {
				str := fmt.Sprint(t[i][j])
				for n, char := range str {
					termbox.SetCell(ox+j*5+1+n, oy+i*2+1, char, fg, bg)
				}
			}
		}
	}
	return termbox.Flush()
}
