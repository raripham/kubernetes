package abstraction

import "fmt"

type SloganSayer interface {
	Slogan()
}

type Campain struct {
	SloganSayer
}

func SaySlogan(s SloganSayer) {
	s.Slogan()
}

type Hillary struct{}

func (h *Hillary) Slogan() {
	fmt.Println("Strong together")
}
