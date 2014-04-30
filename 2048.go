package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Status uint

const (
	Win Status = iota
	Lose
	Add
	Max = 2048
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())

	var table Table
	for {
		st := table.CheckWinAndAdd()
		switch st {
		case Win:
			x, y := termbox.Size()
			//termbox.Clear(termbox.ColorGreen, termbox.ColorYellow)
			str := "Win,use ESC exit"
			strl := len(str)
			for n, c := range str {
				termbox.SetCell(x/2+n-strl/2, y/2, c, 100, 222)
			}
			termbox.Flush()
		case Lose:
			x, y := termbox.Size()
			//termbox.Clear(termbox.ColorRed, termbox.ColorRed)
			str := "Lose,use ESC exit"
			strl := len(str)
			for n, c := range str {
				termbox.SetCell(x/2+n-strl/2, y/2, c, 354, 222)
			}
			termbox.Flush()
		case Add:
			table.Print()
		default:
			fmt.Print("an err input")
		}
		//here get input, only keyarrow is pass,and return keyValue
		key := table.MergeAndReturnKey()
		if key == termbox.KeyEsc {
			return
		}
	}

}

type Table [4][4]int

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
			if t[i][j] > 0 {
				str := fmt.Sprint(t[i][j])
				for n, char := range str {
					termbox.SetCell(j*5+1+n, i*2+1, char, fg, bg)
				}
			}
		}
	}
	return termbox.Flush()
}
func (t *Table) CheckWinAndAdd() Status {
	for _, x := range t {
		for _, y := range x {
			if y >= Max {
				return Win
			}
		}
	}
	i := rand.Intn(len(t))
	j := rand.Intn(len(t))
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if t[i%len(t)][j%len(t)] == 0 {
				t[i%len(t)][j%len(t)] = 2 << (rand.Uint32() % 2)
				return Add
			}
			j++
		}
		i++
	}
	return Lose
}

func (t *Table) MergeAndReturnKey() termbox.Key {
lable:
	ev := termbox.PollEvent()
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			t.MergeUP()
		case termbox.KeyArrowDown:
			t.MergeDwon()
		case termbox.KeyArrowLeft:
			t.MergeLeft()
		case termbox.KeyArrowRight:
			t.MergeRight()
		default:
			fmt.Print("err key")
		}

	case termbox.EventResize:
		t.Print()
		goto lable
	case termbox.EventError:
		panic(ev.Err)
	}
	return ev.Key
}

func (t *Table) MergeUP() {
	tl := len(t)
	for i := 0; i < tl; i++ {
		//去0
		np := tl
		n := 0
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				n++
			}
		}
		//合并
		for x := 0; x < n-1; x++ {
			if t[x][i] == t[x+1][i] {
				t[x][i] *= 2
				t[x+1][i] = 0
				x++
				//	n--
			}
		}
		//再去0
		np = n
		n = 0
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				n++
			}
		}
		if n > np {
			fmt.Print("errrrrrrrrrrrrrrrr")
		}
		//补零
		for x := n; x < tl; x++ {
			t[x][i] = 0
		}
	}
}
func (t *Table) MergeDwon() {
	tl := len(t)
	for i := 0; i < tl; i++ {
		//去0
		np := tl
		n := 0
		for x := 0; x < np; x++ {
			if t[tl-x-1][i] != 0 {
				t[tl-n-1][i] = t[tl-x-1][i]
				n++
			}
		}
		//合并
		for x := 0; x < n-1; x++ {
			if t[tl-x-1][i] == t[tl-x-2][i] {
				t[tl-x-1][i] *= 2
				t[tl-x-2][i] = 0
				x++
				//	n--
			}
		}
		//再去0
		np = n
		n = 0
		for x := 0; x < np; x++ {
			if t[tl-x-1][i] != 0 {
				t[tl-n-1][i] = t[tl-x-1][i]
				n++
			}
		}
		//补零
		for x := n; x < tl; x++ {
			t[tl-x-1][i] = 0
		}
	}
}
func (t *Table) MergeLeft() {
	tl := len(t)
	for i := 0; i < tl; i++ {
		//去0
		np := tl
		n := 0
		for x := 0; x < np; x++ {
			if t[i][x] != 0 {
				t[i][n] = t[i][x]
				n++
			}
		}
		//合并
		for x := 0; x < n-1; x++ {
			if t[i][x] == t[i][x+1] {
				t[i][x] *= 2
				t[i][x+1] = 0
				x++
				//	n--
			}
		}
		//再去0
		np = n
		n = 0
		for x := 0; x < np; x++ {
			if t[i][x] != 0 {
				t[i][n] = t[i][x]
				n++
			}
		}
		//补零
		for x := n; x < tl; x++ {
			t[i][x] = 0
		}
	}
}
func (t *Table) MergeRight() {
	tl := len(t)
	for i := 0; i < tl; i++ {
		//去0
		np := tl
		n := 0
		for x := 0; x < np; x++ {
			if t[i][tl-x-1] != 0 {
				t[i][tl-n-1] = t[i][tl-x-1]
				n++
			}
		}
		//合并
		for x := 0; x < n-1; x++ {
			if t[i][tl-x-1] == t[i][tl-x-2] {
				t[i][tl-x-1] *= 2
				t[i][tl-x-2] = 0
				x++
				//	n--
			}
		}
		//再去0
		np = n
		n = 0
		for x := 0; x < np; x++ {
			if t[i][tl-x-1] != 0 {
				t[i][tl-n-1] = t[i][tl-x-1]
				n++
			}
		}
		//补零
		for x := n; x < tl; x++ {
			t[i][tl-x-1] = 0
		}
	}
}
