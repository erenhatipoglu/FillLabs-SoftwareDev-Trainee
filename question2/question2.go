package main

// generateSequence generates a sequence of numbers based on the input value `n`.
func generateSequence(n int) []int {
	// If `n` is less than or equal to 0, return an empty slice.
	if n <= 0 {
		return []int{}
	}

	// Use a switch statement to determine the sequence to return based on the value of `n`.
	switch {
	case n >= 9: // If `n` is greater than or equal to 9, return the sequence [2, 4, 9].
		return []int{2, 4, 9}
	case n >= 4: // If `n` is greater than or equal to 4 but less than 9, return [2, 4].
		return []int{2, 4}
	case n >= 2: // If `n` is greater than or equal to 2 but less than 4, return [2].
		return []int{2}
	default: // For all other cases, return an empty slice.
		return []int{}
	}
}
