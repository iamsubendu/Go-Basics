//2 types of bit operator(left and right shift)

package main

import "fmt"

func main() {
	a := 10 //1010
	b := 3  //0011

	fmt.Printf("%v\n", a&b)  //and - 0010
	fmt.Printf("%v\n", a|b)  //or	-1011
	fmt.Printf("%v\n", a^b)  //exclusive or -	opp of and -  1001
	fmt.Printf("%v\n", a&^b) //and not	-	oop of or -		0100

	//bit shifty
	a = 8 //2^3
	//00001000

	a = a << 3 //2^3 * 2^3 = 64
	//01000000
	fmt.Println(a)
	a = 8
	a = a >> 3 //2^3 / 2^3 = 2^0 = 1
	//del from end
	//01
	fmt.Println(a)

	x := 42
	fmt.Printf("%d\t%b\n", x, x)
	//42	101010
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y)

	const (
		//kb= 1024
		kb = 1 << 10
		mb = 1024 * kb
		gb = 1024 * mb
	)
	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
