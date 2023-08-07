package utils

import "strconv"

func CheckErrors(tabstr []string) (bool, []int) {
	tabNum := []int{}
	for i := 0; i < len(tabstr); i++ {
		num, err := strconv.Atoi(tabstr[i])
		if err != nil {
			return false, nil
		}
		tabNum = append(tabNum, num)
	}
	for i := 0; i < len(tabNum); i++ {
		for j := 0; j < len(tabNum); j++ {
			if i == j {
				continue
			}
			if tabNum[i] == tabNum[j] {
				return false, nil
			}
		}
	}

	return true, tabNum
}