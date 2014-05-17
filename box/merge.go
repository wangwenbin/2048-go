package box

import (
	"github.com/nsf/termbox-go"
)

func (t *Box) MergeAndReturnKey() termbox.Key {
	var changed bool
Lable:
	changed = false
	ev := termbox.PollEvent()
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			changed = t.MergeUP()
		case termbox.KeyArrowDown:
			changed = t.MergeDwon()
		case termbox.KeyArrowLeft:
			changed = t.MergeLeft()
		case termbox.KeyArrowRight:
			changed = t.MergeRight()
		case termbox.KeyEsc, termbox.KeyEnter:
			changed = true
		default:
			changed = false
			//t.Print(0, 3)
		}
		if !changed {
			goto Lable
		}

	case termbox.EventResize:
		x, y := termbox.Size()
		t.Print(x/2-10, y/2-4)
		goto Lable
	case termbox.EventError:
		panic(ev.Err)
	}
	step++
	return ev.Key
}

func (t *Box) MergeUP() bool {
	tl := len(t)
	changed := false
	notfull := false
	for i := 0; i < tl; i++ {

		np := tl //the last number needed check, first time use len(t).
		n := 0   //count of none 0.

		//clean 0 from top to the last number and move numbers together.
		//imag another t that smaller than this, but covered this.
		//n after "for" is size of the small t,  gives the value of next np.
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				if n != x {
					changed = true
				}
				n++
			}
		}
		if n < tl {
			notfull = true
		}
		np = n
		//mergeup all the number x that are same with its uper one.
		//uper one store 2*x, downer store 0.
		for x := 0; x < np-1; x++ {
			if t[x][i] == t[x+1][i] {
				t[x][i] *= 2
				t[x+1][i] = 0
				Score += t[x][i] * step
				x++
				changed = true
				//	n--
			}
		}
		//clean the new added 0 use the same way.
		n = 0
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				n++
			}
		}
		//cover the unchecked with 0
		for x := n; x < tl; x++ {
			t[x][i] = 0
		}
	}
	return changed || !notfull
}
func (t *Box) MergeDwon() bool {
	t.MirrorV()
	changed := t.MergeUP()
	t.MirrorV()
	return changed
}
func (t *Box) MergeLeft() bool {
	t.Right90()
	changed := t.MergeUP()
	t.Left90()
	return changed
}
func (t *Box) MergeRight() bool {
	t.Left90()
	changed := t.MergeUP()
	t.Right90()
	return changed
}
