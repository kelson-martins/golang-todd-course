package main

import "fmt"

func main() {

	for i := 65; i <= 122; i++ {

		if (i >= 91) && (i <= 96) {
			continue
		}

		fmt.Printf("Integer: %v\tCharacter: %#U\n",i , i)
	}
}