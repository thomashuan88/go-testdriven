package struct_test

import "testing"

type Person struct {
	Name string
	Age  int
}

func (p Person) changeName(n string) { // without * , you cannot change Name
	p.Name = n
}

func TestChangeName(t *testing.T) {
	pp := Person{"huan", 33}

	t.Log(pp)
	pp.Name = "chuan"
	t.Log(pp)
	pp.changeName("heng")

	t.Log(pp)

	yy := Person{"chong", 44}
	t.Log(yy)
	t.Log(pp)

}
