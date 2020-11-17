package main

/*
access all elements of slice without using 'range'
*/

import "fmt"

func main() {

	x := []int{2,5,10,52,22}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}