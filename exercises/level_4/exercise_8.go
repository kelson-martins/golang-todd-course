package main

import "fmt"

/*

1. create a map with a key of TYPE string which is a person's "last_first" name, and avalue of TYPE [string] which stores their favorite things.package level_4
2. store three records in your map.
3. print all of the values, along with their index position in the slice

`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

*/

func main() {

	var myMap = make(map[string][]string)

	myMap["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	myMap["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for name, mapArray := range myMap {
		fmt.Printf("Name: %v\n", name)
		for i, v := range mapArray {
			fmt.Printf("index: %v, %v\n", i, v)
		}
	}
}
