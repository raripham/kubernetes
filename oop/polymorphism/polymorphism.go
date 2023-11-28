package polymorphism

import "fmt"

type SloganSayer interface {
	Slogan()
}

// SaySlogan truyền vào một tham số kiểu SloganSayer
func SaySlogan(sloganSayer SloganSayer) {
	sloganSayer.Slogan()
}

// Hillary thỏa mãn SloganSayer interfa
// bằng việc thực thi function Slogan.
// Vì vậy, Hillary cũng là một kiểu của SloganSayer.
type Hillary struct{}

func (h *Hillary) Slogan() {
	fmt.Println("Stronger together.")
}

// tương tự với struct Trump
type Trump struct{}

func (t *Trump) Slogan() {
	fmt.Println("Make America great again.")
}
