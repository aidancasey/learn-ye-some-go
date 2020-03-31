package main

import (
	"fmt"
)

//func main() {
//
//	var i int
//
//	i = 42
//	var j string
//
//	j = strconv.Itoa(i)
//
//	fmt.Println(j + "fff")
//
//}
func exploreArrays() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := a
	b[7] = 999

	fmt.Println(a)

	// note copying an array makes a new copy
	fmt.Println(b)
}

func main() {

	exploreArrays()

}
