package main

import (
	"fmt"
)

func arrayCopyingGlitches() {

	fmt.Println("*** arrayCopyingGlitches ***")

	a := [...]int{1, 2, 3, 4, 5}

	b := a // here b will point to a seperate copy of the array
	b[1] = 100

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}

func sliceCopyingGlitches() {

	fmt.Println("*** sliceCopyingGlitches ***")
	a := []int{1, 2, 3, 4, 5}

	b := a     // splices share same copy
	b[1] = 100 // here b will point to a seperate copy of the array

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}

func sliceShortcuts() {
	fmt.Println("*** sliceShortcuts ***")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[5:]  // elements 6 ->10
	d := a[:5]  // elements 0 ->5
	e := a[7:8] // elements from index 7 up to but not including index 8

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)

}

func main() {

	//	arrayCopyingGlitches()
	//	sliceCopyingGlitches()
	sliceShortcuts()

}
