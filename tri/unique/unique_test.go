package unique

import "testing"

func TestSimple(t *testing.T) {
	var tabs []int = []int{1, 3, 5, 6, 9, 15, 36, 1}
	var result []int = tri(tabs)
	if result[0] != 1 || result[1] != 1 || result[2] != 3 {
		t.Fail()
	}
}
