package main

import (
	"fmt"
	"os"
	"push-swap/utils"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "" {
		return
	}

	args := os.Args[1:]

	if len(args) == 1 {
		args = strings.Fields(args[0])
	}

	err, a := utils.CheckErrors(args)

	if !err {
		fmt.Println("Error")
		os.Exit(0)
	}
	if len(a) == 1 {
		return
	} else if len(a) == 2 {
		if a[0] > a[1] {
			utils.Sx(a)
			fmt.Println("sa")
		}
		return
	}
	if utils.IsSorted(a) {
		// fmt.Print("Already sorted\n")
		return
	}

	solveFive(a)

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

	cpt := 0
	b := []int{}
	for len(a) != 3 {

		min, indMin := findMin(a)
		for a[0] != min {
			if indMin == 1 {
				a = utils.Sx(a)
				fmt.Println("sa")
				if utils.IsSorted(a) {
					return a
				}
				cpt++
				// fmt.Println(a)

			} else if indMin > (len(a)-1)/2 {
				a = utils.Rrx(a)
				fmt.Println("rra")
				if utils.IsSorted(a) {
					return a
				}
				cpt++
			} else {
				a = utils.Rx(a)
				fmt.Println("ra")
				if utils.IsSorted(a) {
					return a
				}
				cpt++
				// fmt.Println(a)
			}

		}
		b, a = utils.Px(b, a)
		fmt.Println("pb")
		cpt++

	}
	a, _ = solveThree(a)

	for len(b) != 0 {

		a, b = utils.Px(a, b)
		fmt.Println("pa")
		cpt++

	}

	return a
}
