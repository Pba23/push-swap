package utils

func Sx(a []int) []int {
	temp := a[0]
	a[0] = a[1]
	a[1] = temp

	return a
}

// func Sb(b []int) []int {
// 	temp := b[0]
// 	b[0] = b[1]
// 	b[1] = temp

//		return b
//	}
func Ss(a []int, b []int) ([]int, []int) {
	sa := Sx(a)
	sb := Sx(b)

	return sa, sb
}
