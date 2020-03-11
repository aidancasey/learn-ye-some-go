package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello world !")
	fmt.Printf("the current time  is ", time.Now())
	fmt.Println("My favorite number is", rand.Intn(90))
}
