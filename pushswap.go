package main

import (
	"fmt"
	"os"
	"push-swap/utils"
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
	b := []int{}
	for len(a) != 1 {

		min, indMin := findMin(a)
		for a[0] != min {
			a = utils.Rx(a)
			fmt.Println("ra")
		}
		b, a = utils.Px(b, a)
		fmt.Println("pb")

	}
	b = revTab(b)
	for len(b) != 0 {

		a, b = utils.Px(a, b)
		fmt.Println("pb")

	}
	// b = revTab(b)
	fmt.Println(a)
	fmt.Println(b)

	// if len(a) == 3 {
	// 	c := solveThree(a)
	// 	fmt.Println(c)
	// } else if len(a) > 3 {
	// 	n := len(a) - 3
	// 	b := []int{}
	// 	for i := 0; i < n; i++ {
	// 		b, a = utils.Px(b, a)
	// 	}
	// 	fmt.Println(b)
	// 	fmt.Println(a)
	// }

	// n := len(a) / 2

	// min := 0
	// for i := 1; i < len(a); i++ {
	// 	if a[i] < a[min] {
	// 		min = i
	// 	}
	// }
	// a[0], a[min] = a[min], a[0]

	// min := a[0]
	// indMin := 0
	// for i := 1; i < len(a); i++ {
	// 	if a[i] < min {
	// 		min = a[i]
	// 		indMin = i
	// 	}
	// }
	// fmt.Println(indMin)
	// fmt.Println(min)

	// b := []int{}
	// // fmt.Println(a)
	// for i := 0; i < n; i++ {
	// 	b, a = utils.Px(b, a)
	// }
	// b = revTab(b)

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
func revTab(t []int) []int {
	deb := 0
	fin := len(t) - 1
	for deb < fin {
		temp := t[deb]
		t[deb] = t[fin]
		t[fin] = temp
		deb = deb + 1
		fin = fin - 1
	}
	return t
}
func solveThree(a []int) []int {
	if a[0] > a[1] && a[2] > a[1] && a[0] < a[2] {
		a = utils.Sx(a)
		fmt.Println("sa")
	} else if a[0] > a[1] && a[1] > a[2] && a[0] > a[2] {
		a = utils.Sx(a)
		fmt.Println("sa")
		a = utils.Rrx(a)
		fmt.Println("rra")
	} else if a[0] > a[1] && a[1] < a[2] && a[0] > a[2] {
		a = utils.Rx(a)
		fmt.Println("ra")
	} else if a[0] < a[1] && a[1] > a[2] && a[0] < a[2] {
		a = utils.Sx(a)
		fmt.Println("sa")
		a = utils.Rx(a)
		fmt.Println("ra")
	} else {
		a = utils.Rrx(a)
		fmt.Println("rra")
	}

	return a
}
func findMin(a []int) (int, int) {
	min := a[0]
	indMin := 0
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
			indMin = i
		}
	}

	return min, indMin
}
