/*
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/

package main

import "fmt"

const (
	currentYear = iota + 2020
	next1
	next2
	next3
	next4
)

func main() {
	fmt.Println("Next1:", next1)
	fmt.Println("Next2:", next2)
	fmt.Println("Next3:", next3)
	fmt.Println("Next4:", next4)
}
