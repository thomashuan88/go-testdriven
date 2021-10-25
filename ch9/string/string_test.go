package string_test

import (
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // initial value is ""
	s = "hello"
	t.Log(len(s))

	// s[1] = '3' // string 是不可变的byte slice, will kene error if change the slice element

	//s = "\xE4\xB8\xA5" // can store any binary data
	s = "\xE4\xBA\xBB\xFF" // insert a whatever binary data

	t.Log(s)
	t.Log(len(s))

	s = "中" // cannot use ''
	t.Log(len(s))

	c := []rune(s) // use rune the extra unicode from string
	t.Log(len(c))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestUnicode(t *testing.T) {
	sText := "hello 你好"
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	t.Log(textUnquoted)

	t.Logf("%c", 122)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民和国"
	for _, c := range s { // auto to rune
		t.Logf("%[1]c %[1]x", c)
	}
}
