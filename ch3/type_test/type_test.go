package type_test

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {

	var a int32 = 1
	var b int64
	b = int64(a) // have to change type
	var c MyInt

	c = MyInt(b) // have to change type
	t.Log(a, b, c)
	t.Log("maxint64", math.MaxInt64)
	t.Log("maxfloat64", math.MaxFloat64)
	t.Log("maxuint32", math.MaxUint32)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 // go no support pointer calculation
	t.Log(a, aPtr)

	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("s is empty")
	}
}
