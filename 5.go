//if, else, switch

package main

import (
	"fmt"
	"time"
)

func main() {
	num := 2
	if num == 1 {
		fmt.Println("one")
	} else if num == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("none")
	}

	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")

	}

	fmt.Println("When's Friday???")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	case today + 4:
		fmt.Println("In four days.")
	case today + 5:
		fmt.Println("In five days.")
	default:
		fmt.Println("Too far away")
	}
	fmt.Println(time.Now())
}
