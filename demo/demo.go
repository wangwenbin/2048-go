package demo

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/wangwenbin/2048-go/box"
	"math/rand"
	"time"
)

func Go() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())

A:
	b := box.Box{}
	box.Clear()
	for {
		st := b.CheckWinAndAdd()
		x, y := termbox.Size()
		b.Print(x/2-10, y/2-4)
		switch st {
		case box.Win:
			str := "Win!!"
			strl := len(str)
			box.CoverPrintStr(x/2-strl/2, y/2, str, termbox.ColorMagenta, termbox.ColorYellow)
		case box.Lose:
			str := "Lose!!"
			strl := len(str)
			box.CoverPrintStr(x/2-strl/2, y/2, str, termbox.ColorBlack, termbox.ColorRed)
		case box.Add:
		//	b.Print(x/2-10, y/2-4)
		default:
			fmt.Print("Err")
		}
		//here get input, only keyarrow is pass,and return keyValue
		key := b.MergeAndReturnKey()
		if key == termbox.KeyEsc {
			return
		}
		if key == termbox.KeyEnter {
			goto A
		}
	}
}
