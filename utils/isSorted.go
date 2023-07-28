package utils


func IsSorted(a []int) bool {
	length := len(a)
	if length <= 1 {
		return true
	}

	for i := 1; i < length; i++ {
		if a[i] < a[i-1] {
			return false
		}
	}

	return true
}