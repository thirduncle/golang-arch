package main

import "fmt"

type person struct {
	first string
}

// to demonstrate interface usage, we can create another type
// which will also implement the interface
type secretAgent struct {
	person
	ltk bool
}

// attach the speak method to a receiver of type person so 
// that it can also be of TYPE human
func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a secret agent - this is my name", sa.first)
}

// any TYPE that has the methods specified by an interface
// is also of the interface type
// so a value (e.g. p1) can be of more than one types 
// an interface says, if you have these methods, then you are my type

type human interface {
	speak()
}

// Another way to use the interface is by creating a function that takes the type human
func foo(h human) {
	h.speak()
}


func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}

	fmt.Printf("%T\n", p1)

	sa1 := secretAgent{
		person: person{
			first: "Miss Moneypenny",
		},
		ltk: true,
	}

	var x human
	// I can store p1 to x as 
	// x can also store values of type human
	// but because p1 implements the speak method, it is both of type person and type human
	// so it can also be stored in x
	x = p1
	fmt.Printf("%T\n", x)

	// to conclude, a value in Go can be of more than one type
	
	// interfaces are called abstract types and types like struct are called concrete types
	// since x has method speak defined via the interface, we can call speak from x even though 
	// this was attached to person
	x.speak()

	var y human = sa1

	y.speak()

	// so calling speak fromt the abstact type, it goes to the concrete type and find the implementation of this method
	// and call it

	// So we have one interface which is satisfied by both person and secretagent 
	// but their implementations are different - adds flexibility

	fmt.Println("----------------------")
	// We can call function foo both with an explicit human type
	foo(x)
	foo(y)
	// or a person type whose abstract type is human
	foo(p1)
	foo(sa1)
}