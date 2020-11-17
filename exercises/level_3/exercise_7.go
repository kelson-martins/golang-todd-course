package main

/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/

import "fmt"

func main() {
	
	name := "Bond James"

	if name == "Kel Martins" {
		fmt.Println("Kel Martins wrote this")
	} else if name == "Bond James" {
		fmt.Println("Bond James wrote this")
	} else {
		fmt.Println("Solid Snake wrote this")
	}

}