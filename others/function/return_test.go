package return_test

import "testing"

// return with specify variables
func TestFunctionReturn(t *testing.T) {
	n1, n2 := Names()
	t.Log(n1, n2)
}

func Names() (first string, second string) {
	first = "Foo"
	second = "Bar"
	return
}
