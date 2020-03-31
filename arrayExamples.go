package main

import (
	"fmt"
)

func arrayCopyingGlitches() {

	fmt.Println("*** arrayCopyingGlitches ***")

	a := [...]int{1, 2, 3, 4, 5}

	b := a     // same copy of array
	b[1] = 100 // here b will point to a seperate copy of the array

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}

func spliceCopyingGlitches() {

	fmt.Println("*** splictCopyingGlitches ***")
	a := []int{1, 2, 3, 4, 5}

	b := a     // splices share aame copy
	b[1] = 100 // here b will point to a seperate copy of the array

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}

func main() {

	arrayCopyingGlitches()
	spliceCopyingGlitches()

}
