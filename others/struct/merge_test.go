package struct_test

import "testing"

type PP struct {
	Name string
	Age  int
}

func TestPrint(t *testing.T) {

	tt := PP{
		Name: "safasfa",
		Age:  8,
	}

	t.Logf("%+v", tt)

	yy := &PP{
		Name: "uiuyyiu",
		Age:  9,
	}

	t.Logf("%+v", *yy)
}
