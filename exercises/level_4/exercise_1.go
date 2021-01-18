package main

import "fmt"

/*
using a composite literal:
1. crete an ARRAY which holds 5 values of TYPE int
2. assign values to each indedx position
3. range over the array and print the values out
4. using format printing, print out the type of the array
*/
func main() {
	fmt.Println("Level 4, exercise 1")

	var myArray [5]int // by specifying the size [5], I am creating an array

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5

	for idx, num := range myArray {
		fmt.Println("Position", idx, "value", num)
	}

	fmt.Printf("Type of array: %T\n", myArray)
}
