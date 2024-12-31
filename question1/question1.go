package main

import (
	"fmt"     // Package for formatted I/O, such as printing to the console
	"sort"    // Package providing functions to sort slices
	"strings" // Package for string manipulation functions like counting substrings
)

// sortWords sorts words by the number of 'a's and, in case of a tie, by their lengths in descending order.
func sortWords(words []string) []string {
	// Use sort.SliceStable to sort the words slice. Stable sort ensures equal elements retain their order.
	sort.SliceStable(words, func(i, j int) bool {
		// Count the occurrences of the letter 'a' in the i-th and j-th words
		countA1 := strings.Count(words[i], "a")
		countA2 := strings.Count(words[j], "a")

		// If both words have the same count of 'a's, sort them by length (longer words come first)
		if countA1 == countA2 {
			return len(words[i]) > len(words[j])
		}
		// Otherwise, sort by the number of 'a's in descending order
		return countA1 > countA2
	})
	// Return the sorted slice
	return words
}

func main() {
	// Define a slice of strings as the input
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	// Call the sortWords function to sort the input slice
	sorted := sortWords(input)
	// Print the sorted result
	fmt.Println(sorted)
}
