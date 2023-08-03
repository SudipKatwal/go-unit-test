package randomnumber

import "testing"

func TestGenerateRandomNumber_Uniqueness(t *testing.T) {
	// Number of random numbers to generate for the test
	numRandomNumbers := 100

	// Generate the random numbers and store them in a map
	generatedNumbers := make(map[int]bool)
	for i := 0; i < numRandomNumbers; i++ {
		randomNumber := GenerateRandomNumber(0, 1000) // Assuming the range is 0 to 1000
		if generatedNumbers[randomNumber] {
			t.Errorf("Duplicate random number generated: %d", randomNumber)
		}
		generatedNumbers[randomNumber] = true
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	// Test cases with different input ranges
	testCases := []struct {
		min      int
		max      int
		expected bool
	}{
		{0, 100, true},    // Valid range
		{5, 5, true},      // Min and max are the same
		{10, 5, false},    // Invalid range (max < min)
		{-100, 100, true}, // Negative and positive range
	}

	for _, tc := range testCases {
		result := GenerateRandomNumber(tc.min, tc.max)
		if tc.max >= tc.min && (result < tc.min || result > tc.max) {
			t.Errorf("Generated number %d is outside the range [%d, %d]", result, tc.min, tc.max)
		}
	}
}
