/*
Write a program that
1. assigns an int to a variable
2. prints that int in decimal, binary, and hex
3. shifts the bits of that int over 1 position to the left, and assigns that to a variable
4. prints that variable in decimal, binary, and hex
*/

package main

import "fmt"

func main() {
	myVar := 42
	fmt.Printf("Decimal: %d\tBinary: %b\tHexadecimal: %#X\n", myVar, myVar, myVar)
	shiftedVar := myVar << 1
	fmt.Printf("Decimal: %d\tBinary: %b\tHexadecimal: %#X\n", shiftedVar, shiftedVar, shiftedVar)
}
