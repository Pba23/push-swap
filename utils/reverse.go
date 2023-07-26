package utils

func Rrx(a []int) []int {
	last := a[len(a)-1]
	for i := len(a) - 1; i > 0; i-- {
		a[i] = a[i-1]
	}
	a[0] = last
	return a
}

//	func Rrb(b []int) []int {
//		last := b[len(b)-1]
//		for i := len(b) - 1; i > 0; i-- {
//			b[i] = b[i-1]
//		}
//		b[0] = last
//		return b
//	}
func Rrr(a []int, b []int) ([]int, []int) {
	rra := Rrx(a)
	rrb := Rrx(b)

	return rra, rrb
}
