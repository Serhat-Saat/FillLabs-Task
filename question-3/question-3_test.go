package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMostRepeated(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"apple", "pie", "apple", "red", "red", "red"}, "red"},
		{[]string{"dog", "cat", "dog", "dog", "mouse", "cat"}, "dog"},
		{[]string{"1", "2", "3", "1", "4", "5", "1", "6", "2", "2", "2"}, "2"},
		{[]string{"unique"}, "unique"},
		{[]string{}, ""},
	}

	// Run the function for each test case and check the result
	for _, test := range tests {
		result := mostRepeated(test.input)
		fmt.Printf("Input: %v | Expexted: %q, Output: %q \n", test.input, test.expected, result)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Input: %v | Expexted: %q, Output: %q \n", test.input, test.expected, result)
		}
	}
}
