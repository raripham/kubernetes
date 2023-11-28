package encapsulation

import "fmt"

type Encapsulation struct{}

func (e *Encapsulation) Expose() {
	fmt.Println("AHHHH! I'm exposed!")
}

func (e *Encapsulation) hide() {
	fmt.Println("Shhhhhhhhhhhh.. This is super secret")
}

func (e *Encapsulation) Unhide() {
	e.hide()
	fmt.Println("...end")
}
