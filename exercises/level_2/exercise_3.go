/*
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

package main

import "fmt"

const (
	a     = 5
	b int = 20
)

func main() {

	fmt.Println("Untyped Constant value:", a)
	fmt.Println("Typed Constant value:", b)
}
