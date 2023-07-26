package utils

func Px(a []int, b []int) ([]int, []int) {
	a = append(a, b[0])
	b = append(b[:0], b[1:]...)

	return a, b
}

// func Pb(a []int, b []int) ([]int, []int) {
// 	b = append(b, a[0])
// 	a = append(a[:0], a[1:]...)

// 	return a, b
// }
