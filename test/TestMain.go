package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Table [4][4]int

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var t Table
	for i, line := range t {
		for j := range line {
			t[i][j] = i*len(line) + j
		}
	}

	t.Print()
	termbox.PollEvent()

	Transpose(&t).Print()
	termbox.PollEvent()

	Left90(&t).Print()
	termbox.PollEvent()

	Right90(&t).Print()
	termbox.PollEvent()

	MirrorH(&t).Print()
	termbox.PollEvent()

	MirrorV(&t).Print()
	termbox.PollEvent()
}

func (t Table) Print() error {
	fg := termbox.ColorYellow
	bg := termbox.ColorBlack
	termbox.Clear(fg, bg)
	fg = termbox.ColorBlack
	bg = termbox.ColorGreen
	for i := 0; i <= len(t); i++ {
		for x := 0; x < 5*len(t); x++ {
			termbox.SetCell(x, i*2, '-', fg, bg)
		}
		for x := 0; x <= 2*len(t); x++ {
			if x%2 == 0 {
				termbox.SetCell(i*5, x, '+', fg, bg)
			} else {
				termbox.SetCell(i*5, x, '|', fg, bg)
			}
		}
	}
	fg = termbox.ColorYellow
	bg = termbox.ColorBlack
	for i := range t {
		for j := range t[i] {
			str := fmt.Sprint(t[i][j])
			for n, char := range str {
				termbox.SetCell(j*5+1+n, i*2+1, char, fg, bg)
			}
		}
	}
	return termbox.Flush()
}
