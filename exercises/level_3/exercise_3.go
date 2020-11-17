package main

/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/

import "fmt"

func main() {
	bornYear := 1990
	
	for bornYear <= 2020 {
		fmt.Println(bornYear)
		bornYear++
	}
}