package main

import (
	"testing" // Package used for writing and running unit tests
)

// Test_findMostRepeated defines a set of test cases for the findMostRepeated function
func Test_findMostRepeated(t *testing.T) {
	// Define a slice of test cases
	tests := []struct {
		name string   // Name of the test case
		arr  []string // Input slice of strings
		want string   // Expected output from findMostRepeated
	}{
		{
			name: "Example case", // Test case where "red" is the most repeated word
			arr:  []string{"apple", "pie", "apple", "red", "red", "red"},
			want: "red", // Expected output
		},
		{
			name: "Single element", // Test case with a single element in the input slice
			arr:  []string{"onlyone"},
			want: "onlyone", // Expected output is the single element
		},
		{
			name: "All unique elements", // Test case where all elements are unique
			arr:  []string{"a", "b", "c", "d"},
			want: "a", // Expected output is the first element (arbitrarily chosen)
		},
		{
			name: "Empty array", // Test case with an empty input slice
			arr:  []string{},
			want: "", // Expected output is an empty string
		},
		{
			name: "Tie in repetition", // Test case where there is a tie in the maximum repetition count
			arr:  []string{"a", "b", "a", "b", "c"},
			want: "a", // Expected output is the first element with the maximum count
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Run each test case with t.Run for better organization
			// Call findMostRepeated with the test case input and compare the result with the expected output
			if got := findMostRepeated(tt.arr); got != tt.want {
				// Report an error if the result does not match the expected output
				t.Errorf("findMostRepeated() = %v, want %v", got, tt.want)
			}
		})
	}
}
