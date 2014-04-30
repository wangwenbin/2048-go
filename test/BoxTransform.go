// Package this file just provides another way to merger boxes.
// I test this but I didn't use this because I finished write that fuck one
// before this. I still haven't got the best .
package main

func Transpose(t *Table) *Table {
	tn := new(Table)
	for i, line := range t {
		for j, num := range line {
			tn[j][i] = num
		}
	}
	return tn
}
func Right90(t *Table) *Table {
	tn := new(Table)
	for i, line := range t {
		for j, num := range line {
			tn[j][len(t)-i-1] = num
		}
	}
	return tn
}
func Left90(t *Table) *Table {
	tn := new(Table)
	for i, line := range t {
		for j, num := range line {
			tn[len(line)-j-1][i] = num
		}
	}
	return tn
}
func MirrorV(t *Table) *Table {
	tn := new(Table)
	for i, line := range t {
		for j, num := range line {
			tn[len(t)-i-1][j] = num
		}
	}
	return tn
}
func MirrorH(t *Table) *Table {
	tn := new(Table)
	for i, line := range t {
		for j, num := range line {
			tn[i][len(line)-j-1] = num
		}
	}
	return tn
}
