package main

/*
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
*/

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {

		for j := 0; j<= 2; j++ {
			fmt.Printf("%#U\n", i)
		}

	}
}