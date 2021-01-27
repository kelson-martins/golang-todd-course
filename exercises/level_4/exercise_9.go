package main

import "fmt"

/*

Using the code from the previous exercise (level4/exercise8) add a record to the map. Now print the map out using a range loop

*/

func main() {

	var myMap = make(map[string][]string)

	myMap["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	myMap["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}
	myMap["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}

	myMap["new_bond"] = []string{`cars`, `watches`, `villans`}

	for name, mapArray := range myMap {
		fmt.Printf("Name: %v\n", name)
		for i, v := range mapArray {
			fmt.Printf("index: %v, %v\n", i, v)
		}
	}
}
