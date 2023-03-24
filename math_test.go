package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(1, 2)
	if sum != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 3)
	}
}