package box

import (
	"math/rand"
)

type Status uint

const (
	Win Status = iota
	Lose
	Add
	Max = 2048
)

func (t *Box) CheckWinAndAdd() Status {
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
