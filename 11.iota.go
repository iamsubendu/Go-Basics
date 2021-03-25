//The predeclared iota identifier resets to
//0 whenever the word const appears in the source code
//and increments after each const specification:

package main

import (
	"fmt"
)

func main() {
	// 	const (
	// 	C0 = iota
	// 	C1 = iota
	// 	C2 = iota
	// )

	const (
		C0 = iota
		C1
		C2
	)
	fmt.Println(C0, C1, C2) // "0 1 2"

	// const (
	// 	C1 = iota + 1
	// 	C2
	// 	C3
	// )
	// fmt.Println(C1, C2, C3) // "1 2 3"

	const (
		C3 = iota + 3
		_
		C5
		C6
	)
	fmt.Println(C3, C5, C6) // "1 3 4"
}
