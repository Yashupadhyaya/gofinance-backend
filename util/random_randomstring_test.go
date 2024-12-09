package util

import (
	"math/rand"
	"testing"
	"time"
	"unicode/utf8"
)

// Assume that `alphabet` is imported or defined elsewhere in the package
// TODO: Define or import the `alphabet` variable, which should be a non-empty string of characters

func TestRandomString(t *testing.T) {
	t.Parallel() // Allows tests to run in parallel for better performance

	// Seed the random number generator for reproducibility in tests
	rand.Seed(time.Now().UnixNano())

	type testCase struct {
		name     string
		length   int
		validate func(got string) bool
	}

	testCases := []testCase{
		{
			name:   "Generate a Random String of Zero Length",
			length: 0,
			validate: func(got string) bool {
				return got == ""
			},
		},
		{
			name:   "Generate a Random String of Positive Length",
			length: 5,
			validate: func(got string) bool {
				return utf8.RuneCountInString(got) == 5
			},
		},
		{
			name:   "Generate a Random String with Full Alphabet Coverage",
			length: 100,
			validate: func(got string) bool {
				for _, char := range alphabet {
					if !strings.ContainsRune(got, char) {
						return false
					}
				}
				return true
			},
		},
		{
			name:   "Consistency with Seeded Randomness",
			length: 10,
			validate: func(got string) bool {
				rand.Seed(42)
				first := RandomString(10)
				rand.Seed(42)
				second := RandomString(10)
				return first == second
			},
		},
		{
			name:   "Handle Large Input Values",
			length: 1000000,
			validate: func(got string) bool {
				return utf8.RuneCountInString(got) == 1000000
			},
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Run each test case in parallel
			got := RandomString(tc.length)
			if !tc.validate(got) {
				t.Errorf("Test %s failed: unexpected result for length %d, got: %s", tc.name, tc.length, got)
			} else {
				t.Logf("Test %s succeeded", tc.name)
			}
		})
	}
}
