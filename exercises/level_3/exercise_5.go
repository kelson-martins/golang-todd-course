package main

/*
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
*/

import "fmt"

func main() {
	
	for i := 10; i <= 100; i++ {

		fmt.Printf("Remainder of %v modulus 4 = %v\n", i, i % 4)
	}

}