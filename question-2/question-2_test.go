package main

import (
	"fmt"
	"testing"
)

func TestProcessNumberRecursive(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{9, []int{2, 4, 9}},
		{10, []int{2, 5, 10}},
		{15, []int{3, 7, 15}},
		{1, []int{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test for input %d", test.n), func(t *testing.T) {
			stack := &Stack{items: []int{}}
			processNumberRecursive(test.n, stack)

			// To reverse the elements in the stack
			actual := stack.items
			var reversedActual []int
			for i := len(actual) - 1; i >= 0; i-- {
				reversedActual = append(reversedActual, actual[i])
			}

			// Test result comparison
			if len(reversedActual) != len(test.expected) {
				t.Errorf("Test Failed for input: %d\nExpected: %v\nOutput: %v\n", test.n, test.expected, reversedActual)
			} else {
				for i := range reversedActual {
					if reversedActual[i] != test.expected[i] {
						t.Errorf("Test Failed for input: %d\nExpected: %v\nOutput: %v\n", test.n, test.expected, reversedActual)
						break
					}
				}
			}

			// Print if passed successfully
			fmt.Printf("Test Passed for input: %d\nExpected: %v\nOutput: %v\n\n", test.n, test.expected, reversedActual)
		})
	}
}
