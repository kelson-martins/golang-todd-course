package main

/*
Create a program that uses a switch statement with the switch expression specified 
as a variable of TYPE string with the IDENTIFIER “favSport”.
*/

import "fmt"

func main() {
	
	favSport := "Soccer"

	switch favSport {
	case "Swimming":
		fmt.Println("Ultra Marathon")
	case "Voleyball":
		fmt.Println("Beach")
	case "Soccer":
		fmt.Println("Goal")
	default:
		fmt.Println("What a boring person")
	}

}