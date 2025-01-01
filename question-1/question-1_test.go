package main

import (
	"fmt"
	"reflect" //It is used to compare whether two slices are equal or not.
	"testing"
)

// This function tests sortings

func TestSortWords(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			// Normal test data
			[]string{"aaaasd", "a", "aab", "aaabcd", "ef", "csssssssd", "fdz", "kf", "zc", "lklklklklklklkl", "l"},
			[]string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklkl", "csssssssd", "fdz", "ef", "kf", "zc", "l"},
		},
		{
			// Words that contain only "a"
			[]string{"aaa", "aa", "aaaa", "a", "aaaaa"},
			[]string{"aaaaa", "aaaa", "aaa", "aa", "a"},
		},
		{
			// Words without the letter "a"
			[]string{"ef", "csssssssd", "fdz", "kf", "zc", "lklklklklklklkl"},
			[]string{"lklklklklklklkl", "csssssssd", "fdz", "ef", "kf", "zc"},
		},
		{
			// Words with the letter "a" found only once
			[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"},
			[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"},
		},
		{
			// Single-character words
			[]string{"a", "b", "c", "d", "a", "a", "b"},
			[]string{"a", "a", "a", "b", "c", "d", "b"},
		},
		{
			// Empty words
			[]string{"", "a", "ab", "b", ""},
			[]string{"ab", "a", "b", "", ""},
		},
		{
			// Long words with the same number of "a"
			[]string{"aa", "aab", "aaa", "aaaa", "aaaaa", "aaab"},
			[]string{"aaaaa", "aaaa", "aaab", "aaa", "aab", "aa"},
		},
		{
			// Numbers only
			[]string{"2", "5"},
			[]string{"2", "5"},
		},
		{
			// Strings and Number
			[]string{"asdf", "5", "qwrsfa", "asasgasf"},
			[]string{"asasgasf", "qwrsfa", "asdf", "5"},
		},
	}

	fmt.Println("All tests have been completed. Inputs and outputs tested:")
	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %v", test.input), func(t *testing.T) {
			// Print the input for each test case
			fmt.Printf("Input: %v\n", test.input)
			fmt.Printf("Expected Output: %v\n", test.expected)
			result := sortWords(test.input)
			fmt.Printf("Actual Output: %v\n\n", result)
			//Checks if the result is equal to the expected output.
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
