package main

/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/

import "fmt"

func main() {
	
	bornYear := 1990

	for {

		fmt.Println(bornYear)
		bornYear++

		if bornYear == 2021{
			break
		}

	}

}