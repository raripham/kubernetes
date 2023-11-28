package composition

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	CanSwim   bool
}

// Amy được embedded với kiểu Human
// và do đó Amy có thể gọi các phương thức của Human
type Amy struct {
	Human
}

// Alan được embedded với kiểu Human
type Alan struct {
	Human
}

func (h *Human) Name() {

	fmt.Printf("Hello! My name is %v %v", h.FirstName, h.LastName)
}
func (h *Human) Swim() {

	if h.CanSwim == true {
		fmt.Println("I can swim!")
	} else {
		fmt.Println("I can not swim.")
	}
}
