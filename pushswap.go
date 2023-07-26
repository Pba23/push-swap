package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "" {
		return
	}
	args := strings.Fields(os.Args[1])
	err, a := checkErrors(args)
	if !err {
		fmt.Println("Error")
		os.Exit(0)
	}

	// n := len(a) / 2

	// min := 0
	// for i := 1; i < len(a); i++ {
	// 	if a[i] < a[min] {
	// 		min = i
	// 	}
	// }
	// a[0], a[min] = a[min], a[0]

	min := a[0]
	indMin := 0
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
			indMin = i
		}
	}
	fmt.Println(indMin)
	fmt.Println(min)

	// b := []int{}
	// // fmt.Println(a)
	// for i := 0; i < n; i++ {
	// 	b, a = utils.Px(b, a)
	// }
	// fmt.Println(b)
	// fmt.Println(a)

	// b := []int{1}

}
func checkErrors(tabstr []string) (bool, []int) {
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
