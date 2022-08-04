package table_driven

import "testing"

// var lastOccurred := make([]int, 0xffff) // 65535

func lengthOfNonRepeatingSubStr(s string) int {
	// lastOccurred := make(map[rune]int)
	lastOccurred := make([]int, 0xffff) // 65535
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		// if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},
		// Chinense cases
		{"we这里是慕课网", 8},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}
	}

}

//  go test -bench . // benchmark
// go test -bench . -cpuprofile cpu.out
// apt install graphviz
// go tool pprof cpu.out
