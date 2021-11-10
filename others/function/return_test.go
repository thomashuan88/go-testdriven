package return_test

import (
	"fmt"
	"testing"
)

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

func showMapAddress(tt map[string]int) {
	fmt.Printf("..%v\n", tt)
}

func TestMap(t *testing.T) {
	tMap := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Println(tMap)
	t.Logf("--%v", &tMap)
	showMapAddress(tMap)
}
