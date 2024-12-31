package main

import (
	"reflect" // Package used for deep comparison of slices in tests
	"testing" // Package used for writing and running unit tests
)

// Test_generateSequence defines a set of test cases for the generateSequence function
func Test_generateSequence(t *testing.T) {
	// Define a slice of test cases
	tests := []struct {
		name string // Name of the test case
		n    int    // Input value for generateSequence
		want []int  // Expected output slice
	}{
		{"Input 9", 9, []int{2, 4, 9}},  // Test case where n is 9
		{"Input 4", 4, []int{2, 4}},     // Test case where n is 4
		{"Input 2", 2, []int{2}},        // Test case where n is 2
		{"Input 0", 0, []int{}},         // Test case where n is 0 (should return an empty slice)
		{"Negative input", -1, []int{}}, // Test case where n is negative (should return an empty slice)
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Run each test case using t.Run for better organization
			// Call generateSequence with the test case input and compare the result with the expected output
			if got := generateSequence(tt.n); !reflect.DeepEqual(got, tt.want) {
				// Report an error if the result does not match the expected output
				t.Errorf("generateSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
