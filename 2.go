package main

import (
	"fmt"
)

func main() {
	//loops

	//for loop
	i := 1
	for i <= 5 {
		fmt.Println("yes", i)
		i++
	}

	for i = 1; i <= 5; i++ {
		fmt.Println("no", i)
	}

}
