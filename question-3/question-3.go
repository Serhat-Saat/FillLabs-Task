package main

import (
	"fmt"
)

// The function that finds the most repeating element
func mostRepeated(data []string) string {
	if len(data) == 0 {
		return "" // Return empty string for empty list
	}

	frequency := make(map[string]int)

	// Calculate the frequency of each element
	for _, item := range data {
		frequency[item]++
	}

	// Find the most repeating item
	maxCount := 0
	mostRepeatedItem := ""
	for item, count := range frequency {
		if count > maxCount {
			maxCount = count
			mostRepeatedItem = item
		}
	}

	return mostRepeatedItem
}

func main() {
	input := []string{"apple", "pie", "apple", "red", "red", "red"}
	result := mostRepeated(input)
	fmt.Printf("Input: %v -> Output: %q\n", input, result)
}
