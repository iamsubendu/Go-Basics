package main

import (
	"fmt"

	//its package path
	"../CustomPackage/utility"
	myHelper "../CustomPackage/utility/helper"
	//can have a customized name too
)

func main() {
	fmt.Println("Hello")
	fmt.Println(utility.GetName())
	fmt.Println(utility.GetMyName())
	fmt.Println(myHelper.UseName())
}
