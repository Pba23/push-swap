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
	// value, exists := os.LookupEnv("ARG")
	// if exists {
	// 	fmt.Printf("La variable existe et sa valeur est %s\n", value)
	// } else {
	// 	fmt.Println("La variable n'existe pas")
	// }
	args := os.Args[1:]

	if len(args) == 1 {
		args = strings.Fields(args[0])
	}

	//fmt.Println(args)

	err, a := utils.CheckErrors(args)

	if !err {
		fmt.Println("Error")
		os.Exit(0)
	}
	if len(a) == 1 {
		return
	} else if len(a) == 2 {
		if a[0] > a[1] {
			a = utils.Sx(a)
			fmt.Println("sa")
		} 
		return
	}
	if utils.IsSorted(a) {
		fmt.Print("Already sorted\n")
		return
	}

	a = solveFive(a)

	// fmt.Println(b)

}

// func revTab(t []int) []int {
// 	deb := 0
// 	fin := len(t) - 1
// 	for deb < fin {
// 		temp := t[deb]
// 		t[deb] = t[fin]
// 		t[fin] = temp
// 		deb = deb + 1
// 		fin = fin - 1
// 	}
// 	return t
// }
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
	// b = revTab(b)

	// fmt.Print("a = ")
	// fmt.Println(a)
	a, c := solveThree(a)

	// a = revTab(a)
	for len(b) != 0 {

		a, b = utils.Px(a, b)
		fmt.Println("pa")
		cpt++

	}
	// a = revTab(a)
	fmt.Println()
	fmt.Println("Euuuuh", c + cpt)
	return a
}
