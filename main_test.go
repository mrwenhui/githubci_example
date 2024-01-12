package main

import "testing"

func TestAdd(t *testing.T) {
	// ...
	a, b := 3, 5
	expect := 8
	retVal := add(a, b)
	if expect != retVal {
		t.Errorf("Add(%d, %d) return %d, expect %d", a, b, retVal, expect)
	}
}
