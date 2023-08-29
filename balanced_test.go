package main

import "testing"

func TestIsBalanced(t *testing.T) {
	testStrings := []string{"{}()[]", "(({)})", "(()()([])){}"}
	expected := []bool{true, false, true}
	for i, tt := range testStrings {
		result := IsBalanced(tt)
		if result != expected[i] {
			t.Errorf("Incorrect result. Expect %t got %t", expected[i], result)
		}
	}
}
