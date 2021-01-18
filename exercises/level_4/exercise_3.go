package main

import "fmt"

/*
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
    • [42 43 44 45 46]
    • [47 48 49 50 51]
    • [44 45 46 47 48]
    • [43 44 45 46 47]
*/

func main() {

	myArray := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	block1 := myArray[:5]
	block2 := myArray[5:]
	block3 := myArray[2:7]
	block4 := myArray[1:6]

	fmt.Println(block1)
	fmt.Println(block2)
	fmt.Println(block3)
	fmt.Println(block4)
}
