package main

import "fmt"

func main() {

	test1 := monkeyTrouble(true, true)
	test2 := monkeyTrouble(false, false)
	test3 := monkeyTrouble(true, false)

	fmt.Println(test1, test2, test3)
}

func monkeyTrouble(aSmile, bSmile bool) bool {

	if (aSmile && bSmile) || (!aSmile && !bSmile) {
		return true
	}

	return false
}