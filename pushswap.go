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
	if len(a) == 1 {
		fmt.Print("\n⛔ Impossible : One integer entered !\n\n")
		fmt.Println("   Bye...🔌")
		return
	} else if len(a) == 2 {
		if a[0] > a[1] {
			a = utils.Sx(a)
			fmt.Println("sa")
			fmt.Print("\n🎯 sorted in : ")
			fmt.Println(1)
			fmt.Println()
			fmt.Println(a)
		} else {
			fmt.Print("\n💢 already sorted\n\n")
			fmt.Println("   Bye...🔌")
		}
		return
	}
	if isSorted(a) {
		fmt.Print("\n💢 already sorted\n\n")
		fmt.Println("   Bye...🔌")
		return
	}

	a = solveFive(a)

	fmt.Println()
	fmt.Println(a)
	// fmt.Println(b)

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
				// fmt.Println(tabNum[i])
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
func solveThree(a []int) ([]int, int) {
	cpt := 0
	if a[0] > a[1] && a[2] > a[1] && a[0] < a[2] {
		a = utils.Sx(a)
		fmt.Println("sa")
		cpt++
	} else if a[0] > a[1] && a[1] > a[2] && a[0] > a[2] {
		a = utils.Sx(a)
		fmt.Println("sa")
		a = utils.Rrx(a)
		fmt.Println("rra")
		cpt = cpt + 2
	} else if a[0] > a[1] && a[1] < a[2] && a[0] > a[2] {
		a = utils.Rx(a)
		fmt.Println("ra")
		cpt++
	} else if a[0] < a[1] && a[1] > a[2] && a[0] < a[2] {
		a = utils.Sx(a)
		fmt.Println("sa")
		a = utils.Rx(a)
		fmt.Println("ra")
		cpt = cpt + 2
	} else {
		a = utils.Rrx(a)
		fmt.Println("rra")
		cpt++
	}

	return a, cpt
}
func solveFive(a []int) []int {
	// b := []int{}
	// for i := 0; i < 2; i++ {
	// 	b, a = utils.Px(b, a)
	// }
	// a = solveThree(a)
	cpt := 0
	b := []int{}
	for len(a) != 3 {

		min, indMin := findMin(a)
		for a[0] != min {
			if indMin == 1 {
				a = utils.Sx(a)
				fmt.Println("sa")
				cpt++
				// fmt.Println(a)

			} else if indMin > (len(a)-1)/2 {
				a = utils.Rrx(a)
				fmt.Println("rra")
				cpt++
			} else {
				a = utils.Rx(a)
				fmt.Println("ra")
				cpt++
				// fmt.Println(a)
			}

		}
		b, a = utils.Px(b, a)
		fmt.Println("pb")
		cpt++

	}
	b = revTab(b)

	// fmt.Print("a = ")
	// fmt.Println(a)
	a, c := solveThree(a)

	a = revTab(a)
	for len(b) != 0 {

		a, b = utils.Px(a, b)
		fmt.Println("pa")
		cpt++

	}
	a = revTab(a)
	fmt.Print("\n🎯 sorted in : ")
	fmt.Println(c + cpt)
	return a
}
func isSorted(a []int) bool {
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
