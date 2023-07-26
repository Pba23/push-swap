package utils

func Rx(a []int) []int {
	// a = append(a[1:], a[0])
	first := a[0]
	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = first

	return a
}

// func Rb(b []int) []int {
// 	// b = append(b[1:], b[0])
// 	first := b[0]
// 	for i := 0; i < len(b)-1; i++ {
// 		b[i] = b[i+1]
// 	}
// 	b[len(b)-1] = first

//		return b
//	}
func Rr(a []int, b []int) ([]int, []int) {
	ra := Rx(a)
	rb := Rx(b)

	return ra, rb
}
