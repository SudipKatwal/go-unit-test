package datetime

import (
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	// Test cases with different years
	testCases := []struct {
		year     int
		expected bool
	}{
		{2000, true},  // Divisible by 4, 100, and 400
		{2020, true},  // Divisible by 4 and 100 but not 400
		{2024, true},  // Divisible by 4 but not 100
		{2021, false}, // Not divisible by 4
		{1900, false}, // Divisible by 4 and 100 but not 400
		{1800, false}, // Divisible by 4 but not 100
		{1600, true},  // Divisible by 4, 100, and 400
	}

	for _, tc := range testCases {
		result := IsLeapYear(tc.year)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v for year %d", tc.expected, result, tc.year)
		}
	}
}

func TestConvertToTimeZone(t *testing.T) {
	// Test cases with different time zones
	testCases := []struct {
		utcTime  time.Time
		timeZone string
		expected time.Time
		isError  bool
	}{
		{
			utcTime:  time.Date(2023, time.August, 3, 12, 34, 56, 0, time.UTC),
			timeZone: "America/New_York",
			expected: time.Date(2023, time.August, 3, 8, 34, 56, 0, time.FixedZone("EDT", -4*60*60)),
			isError:  false,
		},
		{
			utcTime:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			timeZone: "Asia/Tokyo",
			expected: time.Date(2023, time.January, 1, 9, 0, 0, 0, time.FixedZone("JST", 9*60*60)),
			isError:  false,
		},
		// Add more test cases here for different time zones and dates
		{
			utcTime:  time.Date(2023, time.July, 15, 10, 0, 0, 0, time.UTC),
			timeZone: "Invalid_TimeZone",
			expected: time.Time{},
			isError:  true,
		},
	}

	for _, tc := range testCases {
		result, err := ConvertToTimeZone(tc.utcTime, tc.timeZone)

		if tc.isError && err == nil {
			t.Errorf("Expected error for time zone %s, but got none", tc.timeZone)
		} else if !tc.isError && err != nil {
			t.Errorf("Unexpected error for time zone %s: %v", tc.timeZone, err)
		} else if !tc.isError && !result.Equal(tc.expected) {
			t.Errorf("Expected %v, but got %v for time zone %s", tc.expected, result, tc.timeZone)
		}
	}
}
