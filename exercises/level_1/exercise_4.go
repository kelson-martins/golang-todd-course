/* 
FYI - nice documentation and new terminology “underlying type”
https://golang.org/ref/spec#Types 
For this exercise

1. Create your own type. Have the underlying type be an int.
2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
3. in func main
	a. print out the value of the variable “x”
	b. print out the type of the variable “x”
	c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	d. print out the value of the variable “x”
*/

package main

import "fmt"

type kekel int

var x kekel

func main() {	
	fmt.Println("Value of x:",x)

	xType := fmt.Sprintf("%T",x)
	fmt.Println("Type x:",xType)

	x = 42
	fmt.Println("New value of x:",x)
}