/*
1. using a composite literal, create a slice of type int
2. assign 10 values
3. range over the slice and print the values
4. print the type of the array
*/

package main

import "fmt"

func main() {

	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for idx, value := range myArray {
		fmt.Println(idx, value)
	}

	fmt.Printf("Slice type: %T\n", myArray)
}
