package main

import (
	"reflect" // Package used for deep comparison of complex data types, like slices, in tests
	"testing" // Package used for writing and executing unit tests
)

// Test_sortWords defines a set of test cases for the sortWords function
func Test_sortWords(t *testing.T) {
	// Define a structure to hold the input arguments for the sortWords function
	type args struct {
		words []string // A slice of strings representing the words to be sorted
	}
	// Define a slice of test cases
	tests := []struct {
		name string   // Name of the test case
		args args     // Input arguments for the test case
		want []string // Expected output of the sortWords function
	}{
		{
			name: "Example case", // Test case to verify sorting based on 'a' count and lexicographical order
			args: args{
				words: []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
			},
			want: []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"},
		},
		{
			name: "All same 'a' counts", // Test case where all words have the same count of 'a'
			args: args{
				words: []string{"abc", "xyz", "kqa"},
			},
			want: []string{"abc", "kqa", "xyz"}, // Expected order is based on lexicographical order
		},
		{
			name: "No 'a' characters", // Test case where none of the words contain the letter 'a'
			args: args{
				words: []string{"zzz", "yyy", "www"},
			},
			want: []string{"zzz", "yyy", "www"}, // Order remains as input since there are no 'a's to compare
		},
		{
			name: "Empty input", // Test case with an empty input slice
			args: args{
				words: []string{}, // No words to sort
			},
			want: []string{}, // Expected output is also an empty slice
		},
	}
	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Run each test case using t.Run for better organization
			// Call sortWords with the test case input and compare the result with the expected output
			if got := sortWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				// Report an error if the result does not match the expected output
				t.Errorf("sortWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
