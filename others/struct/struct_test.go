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

type cc struct {
	Id int
}

func TestPointerStruct(t *testing.T) {
	ii := cc{Id: 44}
	yy := cc{Id: 66}
	changval(&ii)
	t.Log(ii)
	t.Logf("ii is %p, yy is %p", &ii, &yy)
	// ii.id = 33
}

func changval(v *cc) {
	v.Id = 55
}
