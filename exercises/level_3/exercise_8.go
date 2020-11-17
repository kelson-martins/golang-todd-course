package main

/*
Create a program that uses a switch statement with no switch expression specified.
*/

import "fmt"

func main() {
	
	name := "Solid Snake"

	switch {
	case name == "Bond James":
		fmt.Println("On theaters April 2th 2021")
	case name == "Kel Martins":
		fmt.Println("On his way to become Jedi level 3")
	case name == "Solid Snake":
		fmt.Println("Snaaaaake")
	default:
		fmt.Println("default case")
	}

}