package main

import (
	"fmt" // Package for formatted I/O, such as printing to the console
)

// findMostRepeated finds the most repeated string in an array.
func findMostRepeated(arr []string) string {
	// If the input array is empty, return an empty string
	if len(arr) == 0 {
		return ""
	}

	// Create a map to count the occurrences of each string
	counts := make(map[string]int)
	var mostRepeated string // Variable to store the string with the highest repetition
	maxCount := 0           // Variable to track the maximum count of any string

	// Iterate over the array
	for _, word := range arr {
		// Increment the count for the current word in the map
		counts[word]++
		// If the current word's count exceeds the maxCount, update maxCount and mostRepeated
		if counts[word] > maxCount {
			maxCount = counts[word]
			mostRepeated = word
		}
	}

	// Return the most repeated string
	return mostRepeated
}

func main() {
	// Input data: a slice of strings
	data := []string{"apple", "pie", "apple", "red", "red", "red"}
	// Call findMostRepeated and print the result
	fmt.Println("Most Repeated:", findMostRepeated(data))
}
