//use of functions
//their are functions available in the packages
//and also we can create functions too

package main

import (
	"fmt"
)

func add(x int, y int) int {
	//func add(x, y int) int => we can also write like this
	s := x + y
	return s
}

func calc(x, y int) (int, int) {
	z := x - y
	s := x / y
	return z, s
}

// func calc(x, y int) (out1, out2 int) {
// 	z := x - y
// 	s := x / y
// 	return
//by declaring what we want to return in the initial...we dont have to mention the return type
// }

func main() {

	num1 := 5
	num2 := 7
	result := add(num1, num2)
	fmt.Println(result)
	yono1, yono2 := calc(num1, num2)
	fmt.Println(yono1, yono2)

}

//for a function name if we define the function name starting with small letters it won't be available
//outside the current package
//this is very very important
