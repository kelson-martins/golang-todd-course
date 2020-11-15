/*
Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
42
“James Bond”
true
Now print the values stored in those variables using 
a single print statement
multiple print statements
*/

package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	z := true

	printSingle(x,y,z)
	printMultiple(x,y,z)
}

func printSingle(x int, y string, z bool) {
	fmt.Println("Printing variables on sigle statement:", x, y, z)
}

func printMultiple(x int, y string, z bool) {
	fmt.Println("Printing multiple statements:")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}