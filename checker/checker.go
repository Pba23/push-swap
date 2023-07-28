package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/utils"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanLines)

	

	var Args []string
	tabInt := os.Args[1:]

	if len(tabInt) == 0 {
		return
	}

	tabInt = strings.Fields(tabInt[0])

	str := ""

	for scanner.Scan() {
		
		str +=  scanner.Text() + " "
	}
		Args = strings.Fields(str)
		//fmt.Println(Args)

	if !IsValidInt(tabInt) {
		//fmt.Println("1")
		return
	}

	if !checkErrors(Args) {
		//fmt.Println("2")
		return
	}

	if !CheckValidity(Args, tabInt) {
		fmt.Println("KO")
		return
	}

	fmt.Println("OK")
}

func checkErrors(tab []string) bool {
	for _, val := range tab {
		if val != "pa" && val != "pb" && val != "sa" && val != "sb" && val != "ss" && val != "ra" && val != "rb" && val != "rr" && val != "rra" && val != "rrb" && val != "rrr" && val != "\n" {
			//fmt.Println("val")
			return false

		}
	}

	if tab[0] == "\n" {
		//fmt.Println("val")
		return false
	}


	return true
}

func IsValidInt(tab []string) bool {
	for _, val := range tab {
		_, err := strconv.Atoi(val)

		if err != nil {
			
			return false
		}
	}

	return true
}

func CheckValidity(instructions, tobeSorted []string) (bool) {
	stackA := TransformTab(tobeSorted)

	var stackB []int
	for _, val := range instructions {
		switch val {
		case "pa" : 
			stackA, stackB = utils.Px(stackA, stackB)
		case "pb" :
			stackB, stackA = utils.Px(stackB, stackA)
		case "sa" :
			stackA = utils.Sx(stackA)
		case "sb" :
			stackB = utils.Sx(stackB)
		case "ss" : 
			stackA, stackB = utils.Ss(stackA, stackB)
		case "ra" : 
			stackA  = utils.Rx(stackA)
		case "rb" : 
			stackB  = utils.Rx(stackB)
		case "rr" : 
			stackA, stackB = utils.Rr(stackA, stackB)
		case "rra" : 
			stackA  = utils.Rrx(stackA)
		case "rrb" : 
			stackB  = utils.Rrx(stackB)
		case "rrr" : 
			stackA, stackB = utils.Rrr(stackA, stackB)
		}
		fmt.Println("sA", stackA)
		fmt.Println("Sb", stackB)
	}
	if utils.IsSorted(stackA) {
		return true
	}

	return false
}

// pb\nsa\npb\nsa\nrra\npa\npa
// pb\nra\npb\nra\nsa\nra\npa\npa\n

func TransformTab(tab []string) []int {
	var tabInt []int

	for _, val := range tab {
		valInt, _ := strconv.Atoi(val)

		tabInt = append(tabInt, valInt)
	}

	return tabInt
}
