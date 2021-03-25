package main

import "fmt"

type Kloudone interface {
	name() string
	amount() float64
}

type trainee struct {
	fname   string
	balance float64
}

type instructor struct {
	fname   string
	balance float64
}

func (t trainee) name() string {
	//ob
	return t.fname
}

func (i instructor) name() string {
	return i.fname
}

func (t trainee) amount() float64 {
	return t.balance
}

func (i instructor) amount() float64 {
	return i.balance
}

func measure(k Kloudone) {
	fmt.Println(k)
	fmt.Println(k.name())
	fmt.Println(k.amount())
}

func main() {
	t1 := trainee{fname: "Subendu", balance: 9999}
	t2 := instructor{fname: "Ankit", balance: 9999}

	measure(t1)
	measure(t2)
}
