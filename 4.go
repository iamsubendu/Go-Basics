//package math

package main

import (
	"fmt"
	"math"
)

func main() {
	// num:= 9
	//sqrt works with float no not with int
	var num float64 = 12
	result := math.Sqrt(num)
	fmt.Println((result))
	//to print a format
	fmt.Printf("The output is %f thankyou", result)
	//in double quotes we have to write placeholder
	//for float we can use %f/%g
	fmt.Println("\n")
	fmt.Printf("The output is %.2f thankyou", result)
	fmt.Println("\n")
	fmt.Printf("The output is %.2g thankyou", result)
	fmt.Println("\n")

	//for round off result we can use
	//math.Round(result)
	//we can use ceil that will give upward value
	//math.Ceil(result)
	//or we can use floor that will give bottom value
	//math.Floor(result)
}
