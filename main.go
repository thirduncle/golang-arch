package main

import "fmt"

type person struct {
	first string
}

// attach the speak method to a receiver of type person so 
// that it can also be of TYPE human
func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

// any TYPE that has the methods specified by an interface
// is also of the interface type
// so a value (e.g. p1) can be of more than one types 
// an interface says, if you have these methods, then you are my type

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	var x human
	// I can store p1 to x as 
	// x can also store values of type human
	// but because p1 implements the speak method, it is both of type person and type human
	// so it can also be stored in x
	x = p1
	fmt.Printf("%T\n", x)

	// to conclude, a value in Go can be of more than one type
}