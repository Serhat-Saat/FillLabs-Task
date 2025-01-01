package main

import (
	"fmt"
	"sort"
	"strings"
)

// That function sorting the words
func sortWords(words []string) []string {
	//sort.SliceStable function performs a stable order on the given slice, thus preserving the original order of equal elements.
	sort.SliceStable(words, func(i, j int) bool {
		countA1 := strings.Count(words[i], "a")
		countA2 := strings.Count(words[j], "a")

		// Descending order by number of characters "a"
		if countA1 != countA2 {
			return countA1 > countA2
		}

		// If the number "a" is equal, the order descends by word length
		return len(words[i]) > len(words[j])
	})

	return words
}

func main() {
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "csssssssd", "fdz", "kf", "zc", "lklklklklklklkl", "l"}
	sortedWords := sortWords(input)
	fmt.Println("Words To Sort:", input)
	fmt.Println("Sorted Words:", sortedWords)
}
