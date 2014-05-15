// Package this file just provides another way to merger boxes.
// I test this but I didn't use this because I finished write that fuck on
// before this. I still haven't got the best .
package box

func (t *Box) Transpose() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[j][i] = num
		}
	}
	*t = *tn
}
func (t *Box) Right90() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[j][len(t)-i-1] = num
		}
	}
	*t = *tn
}
func (t *Box) Left90() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[len(line)-j-1][i] = num
		}
	}
	*t = *tn
}
func (t *Box) MirrorV() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[len(t)-i-1][j] = num
		}
	}
	*t = *tn
}
func (t *Box) MirrorH() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[i][len(line)-j-1] = num
		}
	}
	*t = *tn
}
