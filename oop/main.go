package main

import (
	"fmt"

	"github.com/raripham/kubernetes/oop/abstraction"
	"github.com/raripham/kubernetes/oop/encapsulation"
	"github.com/raripham/kubernetes/oop/polymorphism"
)

type someStruc struct {
	Name string
}

func (s *someStruc) foo(Name string) {
	s.Name = Name
}

func main() {
	someStruc := new(someStruc)

	someStruc.foo("abs")
	fmt.Printf(someStruc.Name)

	someStruc.Name = "agggg"
	fmt.Printf(someStruc.Name)

	e := encapsulation.Encapsulation{}

	e.Expose()
	e.Unhide()

	hillary := new(polymorphism.Hillary)
	hillary.Slogan()
	polymorphism.SaySlogan(hillary)

	trump := new(polymorphism.Trump)
	polymorphism.SaySlogan(trump)

	hillary2 := new(abstraction.Hillary)
	h := abstraction.Campain{hillary2}

	h.Slogan()
	abstraction.SaySlogan(hillary2)
	abstraction.SaySlogan(h)
}
