//%d for decimal
//\t for tab space in string
//%b for binary value
//\n for new line
//%v for variable

//for anything we have to call functions

// Print("Hello World")
//only this wont work cuz all things are in packages so we have to import the package

package main

//this is also imp cuz by this anyone can use the packages

import "fmt"

// fmt.Print ("Hello World")
//only this wont work still...so we have to use a execution point that is main()

func main() {
	fmt.Print("Hello World \n")

	fmt.Print(2 + 3)
	fmt.Print("\n")
	//variables
	var num = 2
	fmt.Print(num)
	fmt.Print("\n")

	//uint supports positive no and
	//int supports the whole range
	//float
	//bool
	//string
	//rune
	//const

	//by default value for any var is 0

	var num1, num2 int
	num1, num2 = 2, 3
	var result = num1 + num2
	fmt.Print(result)
	fmt.Print("\n")
	fmt.Printf("%v, %T\n", result, result)

	x := 9 //shorthand of (var result = 9)
	fmt.Print(x)
	fmt.Print("\n")
	fmt.Printf("%v, %T\n", x*x, x)

	const fool = 9
	// fool=7    //as u uncomment this you will get error cuz u cant change constant values
	fmt.Print(fool)
	fmt.Print("\n")

}
