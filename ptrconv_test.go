package ptrconv

import "testing"

func TestIntSliceString(t *testing.T) {
	if result := IntSliceString([]int{1, 2, 3, 4}); result != `[1, 2, 3, 4]` {
		t.Fatalf("expected output: %s, got: %s", `[1,2,3,4]`, result)
	}
}
