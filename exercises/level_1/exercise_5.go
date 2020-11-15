/*
Building on the code from the previous exercise
1. at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
eg:

type hotdog int
var x hotdog
var y int

2. in func main
	a. this should already be done
		1. print out the value of the variable “x”
		2. print out the type of the variable “x”
		3. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
		4. print out the value of the variable “x”
	b. now do this
		1. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
		2. then use the “=” operator to ASSIGN that value to “y”
		3. print out the value stored in “y”
		4. print out the type of “y”

*/

package main

import "fmt"

type kekel int

var x kekel
var y int

func main() {

	fmt.Println("Value of x:", x)
	xType := fmt.Sprintf("%T", x)
	fmt.Println("Type x:", xType)
	x = 42
	fmt.Println("New value of x:", x)

	xConverted := int(x)
	y = xConverted
	fmt.Println("value of y (converted x):", y)
	fmt.Printf("%s %T\n", "new type of y:", y)
}
