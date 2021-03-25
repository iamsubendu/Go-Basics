//In Go language, a channel is a medium through which a
//goroutine communicates with another goroutine and this
//communication is lock-free. Or in other words, a
//channel is a technique which allows to let one goroutine
//to send data to another goroutine.

//Channels are a typed conduit through which you
//can send and receive values with the channel operator, <-.
//(The data flows in the direction of the arrow.)

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	//elements before median

	go sum(s[len(s)/2:], c)
	//elements after median

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
