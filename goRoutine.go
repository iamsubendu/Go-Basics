//lightweight thread and runs concurrently

//speacial func => go concurrency => parallel function which runs with main program
//program which runs concurrently is called go Routine

//instead of using thread we use goRoutine.

//A Goroutine is a function or method which executes independently and
//simultaneously in connection with any other Goroutines present in your
//program. Or in other words, every concurrently executing activity in Go
//language is known as a Goroutines. You can consider a Goroutine like a
//light weighted thread.

package main

import (
	"fmt"
	"time"
)

// func display(str string) {
// 	for w := 0; w < 6; w++ {
// 		fmt.Println(str)
// 	}
// }

// we simply create a display() function and then call this function in two different ways
// first one is a Goroutine, i.e. go display(“Welcome”) and another one is a normal function,
// i.e. display(“GeeksforGeeks”). But there is a problem, it only displays the result of the
// normal function that does not display the result of Goroutine because when a new Goroutine executed,
// the Goroutine call return immediately. The control does not wait for Goroutine to complete their
// execution just like normal function they always move forward to the next line after the Goroutine
// call and ignores the value returned by the Goroutine.

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

// We added the Sleep() method in our program which makes the main Goroutine sleeps for 1 second
// in between 1-second the new Goroutine executes, displays “welcome” on the screen, and then terminate
// after 1-second main Goroutine re-schedule and perform its operation. This process continues until
// the value of the z<6 after that the main Goroutine terminates. Here, both Goroutine and the normal function work concurrently.

func main() {

	// Calling Goroutine
	go display("Welcome")

	// Calling normal function
	display("GeeksforGeeks")
}
